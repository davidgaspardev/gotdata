package mssql

import (
	"fmt"
	"testing"
	"time"

	"github.com/davidgaspardev/gotdata/helpers"
)

func Test_WriteInMultithread(t *testing.T) {

	gotdb := GetInstance()
	threadResult := make(chan uint)
	threadNumber := 1000

	if err := gotdb.Exec(`
		IF NOT EXISTS (SELECT * FROM [sysobjects] WHERE [name] = 'PLC_pulses')
		BEGIN
			CREATE TABLE [PLC_pulses] (
				[id]       INTEGER PRIMARY KEY IDENTITY(1,1),
				[mac_code] VARCHAR(16) NOT NULL,
				[date]     DATETIME    NOT NULL,
				[counter]  INTEGER     NOT NULL
			)
		END
	`); err != nil {
		t.Log(err)
		t.FailNow()
	}

	for i := 0; i < threadNumber; i++ {
		go func(id uint) {
			loop := 100

			for j := 0; j < loop; j++ {
				now := time.Now()

				data := map[string]interface{}{
					"mac_code": fmt.Sprintf("INJ%d", j),
					"date":     now.Format("2006-01-02 15:04:05"),
					"counter":  j,
				}

				if err := gotdb.Write("PLC_pulses", data); err != nil {
					t.Error(err)
				}
			}

			threadResult <- id
		}(uint(i))
	}

	for i := 0; i < threadNumber; i++ {
		threadId := <-threadResult

		t.Logf("%dº thread finished", threadId)
	}

	if err := gotdb.Close(); err != nil {
		t.Error(err)
	}
}

func Test_ReadInMultithread(t *testing.T) {
	gotdb := GetInstance()
	threadResult := make(chan uint)
	threadNumber := 1000

	for i := 0; i < threadNumber; i++ {
		go func(id uint) {
			filter := helpers.Filter{
				Page: (id + 1),
			}

			data, err := gotdb.Read("PLC_pulses", []string{"mac_code", "date", "counter"}, &filter)
			if err != nil {
				t.Error(err)
			}

			for j := 0; j < len(data); j++ {
				t.Logf("%dº - Data: %+v\n", j, data[j])
			}

			threadResult <- id
		}(uint(i))
	}

	for i := 0; i < threadNumber; i++ {
		threadId := <-threadResult

		t.Logf("%dº thread finished", threadId)
	}

	gotdb.Close()
}

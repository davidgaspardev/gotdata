package mssql

import (
	"fmt"
	"testing"
	"time"
)

func Test_WriteInMultithread(t *testing.T) {

	gotdb := GetInstance()
	threadResult := make(chan uint)
	threadNumber := 1000

	gotdb.Exec(`
		IF NOT EXISTS (SELECT * FROM [sysobjects] WHERE [name] = 'PLC_pulses')
		BEGIN
			CREATE TABLE [PLC_pulses] (
				[id]       INTEGER PRIMARY KEY IDENTITY(1,1),
				[mac_code] VARCHAR(16) NOT NULL,
				[date]     DATETIME    NOT NULL,
				[counter]  INTEGER     NOT NULL
			)
		END
	`)

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

	gotdb.Close()
}

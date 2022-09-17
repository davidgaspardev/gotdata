package main

import (
	"fmt"
	"gotdata"
	"time"
)

var gotdb gotdata.Gotdata

func main() {
	gotdb = gotdata.GetGotdata()

	if err := gotdb.Exec(`
		IF NOT EXISTS (SELECT * FROM [sysobjects] WHERE [name] = 'PLC_pulses')
		BEGIN
			CREATE TABLE [PLC_pulses] (
				[id]       INTEGER PRIMARY KEY IDENTITY(1,1),
				[mac_code] VARCHAR(16) NOT NULL,
				[date]     DATETIME    NOT NULL,
				[counter]  INTEGER     NOT NULL
			);
		END

		DELETE [PLC_pulses];
	`); err != nil {
		panic(err)
	}

	writeDataInDatabase()
	readDataInDatabase()
	updateDataInDatabase() // Show rows updated'
	deleteDataInDatabase()

	gotdb.Close()
}

func writeDataInDatabase() {
	dataMap := map[string]interface{}{
		"mac_code": "INJ001",
		"date":     time.Now().Format("2006-01-02 15:04:05"),
		"counter":  8,
	}

	if err := gotdb.Write("PLC_pulses", dataMap); err != nil {
		panic(err)
	}
}

func readDataInDatabase() {
	data, err := gotdb.Read("PLC_pulses", []string{"mac_code", "counter"}, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(data) // [map[counter:8 mac_code:INJ001]]
}

func updateDataInDatabase() {
	dataMap := map[string]interface{}{
		"counter": 16,
	}

	if err := gotdb.Update("PLC_pulses", dataMap, nil); err != nil {
		panic(err)
	}
}

func deleteDataInDatabase() {
	if err := gotdb.Delete("PLC_pulses", nil); err != nil {
		panic(err)
	}
}

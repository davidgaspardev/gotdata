package tools

import "testing"

func Test_SelectFrom(t *testing.T) {
	// Use this attributes
	attributes := []string{"id", "mac_code", "date", "counter"}

	tSqlBuilder := TSqlBuilder{}

	tSqlBuilder.Select(attributes).From("MAC_machines")

	if result := tSqlBuilder.Done(); result != "SELECT id,mac_code,date,counter FROM [MAC_machines];" {
		t.Error(result)
	} else {
		t.Log(result)
	}
}

func Test_SelectColumnsFrom(t *testing.T) {
	// Use this attributes
	attributes := []string{"code", "name", "type", "stopFactor", "status"}

	tSqlBuilder := TSqlBuilder{}

	tSqlBuilder.SelectColumns(attributes).From("MAC_machines")

	if result := tSqlBuilder.Done(); result != "SELECT [code],[name],[type],[stopFactor],[status] FROM [MAC_machines];" {
		t.Error(result)
	} else {
		t.Log(result)
	}
}

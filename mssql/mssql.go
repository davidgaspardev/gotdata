package mssql

import (
	"database/sql"
	"fmt"

	"github.com/davidgaspardev/golog"
)

type _DatabaseMssql struct {
	logger bool
	db     *sql.DB
}

var instance *_DatabaseMssql

func init() {
	// Check if has connection
	if !hasConnectionData() {
		err := fmt.Errorf("don't has connection data")
		golog.Error("Mssql", err)
		panic(err)
	}

	// Make connection to the database
	db, err := connectDatabase()
	if err != nil {
		golog.Error("Mssql", err)
		panic(err)
	}

	// Configure database connection for better performance
	configDatabaseConnection(db)

	// Alloc instance
	instance = &_DatabaseMssql{
		db: db,
	}
}

func connectDatabase() (db *sql.DB, err error) {
	// Create pool connection
	db, err = sql.Open("sqlserver", createConnectionUrl())

	return
}

func GetInstance() *_DatabaseMssql {
	return instance
}

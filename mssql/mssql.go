package mssql

import (
	"database/sql"
	"fmt"

	"github.com/davidgaspardev/golog"
	_ "github.com/denisenkom/go-mssqldb"
)

type _DatabaseMssql struct {
	logger bool
	db     *sql.DB
}

var instance *_DatabaseMssql

func createInstance() {
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
	if instance == nil {
		createInstance()
	}

	return instance
}

func (mssql *_DatabaseMssql) Restart() error {
	if err := mssql.db.Close(); err != nil {
		return err
	}

	// Delete instance
	instance = nil

	createInstance()

	return nil
}

func (mssql *_DatabaseMssql) Close() error {
	return mssql.db.Close()
}

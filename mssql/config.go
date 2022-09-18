package mssql

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
)

// Get all the data to make tha connection
func getConnectionData() (username string, password string, host string, port string, instance string, database string) {
	username = os.Getenv("MSSQL_USERNAME")
	password = os.Getenv("MSSQL_PASSWORD")
	host = os.Getenv("MSSQL_HOST")
	port = os.Getenv("MSSQL_PORT")
	instance = os.Getenv("MSSQL_INSTANCE")
	database = os.Getenv("MSSQL_DATABASE")

	return
}

// Check if exists connection
func hasConnectionData() bool {
	var username,
		password,
		host,
		port,
		instance,
		database = getConnectionData()

	return (username != "" &&
		password != "" &&
		host != "" &&
		port != "" &&
		instance != "" &&
		database != "")
}

func createConnectionUrl() string {
	var username,
		password,
		host,
		port,
		instance,
		database = getConnectionData()

	query := url.Values{}
	query.Add("database", database)

	url := &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(username, password),
		Host:     fmt.Sprintf("%s:%s", host, port),
		Path:     instance,
		RawQuery: query.Encode(),
	}

	return url.String()
}

func configDatabaseConnection(db *sql.DB) {
	db.SetMaxOpenConns(8)
	db.SetMaxIdleConns(4)
}

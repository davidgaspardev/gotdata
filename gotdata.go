package gotdata

import (
	"github.com/davidgaspardev/gotdata/helpers"
	"github.com/davidgaspardev/gotdata/mssql"
)

type Filter = helpers.Filter
type Where = helpers.Where

type Gotdata interface {
	// Debug
	SetLogger(showLog bool)

	// Execute T-SQL statement
	Exec(tSql string) error
	// Query data with T-SQL statement
	Query(tSql string) ([]map[string]interface{}, error)
	// Count rows from table
	Count(tableName string) (uint, error)
	// Insert data in the table
	Write(tableName string, data map[string]interface{}) error
	// Select columns (attributes) from table with optional filter
	Read(tableName string, attributes []string, filter *Filter) ([]map[string]interface{}, error)
	// Update row(s) in the table with optional filter
	Update(tableName string, data map[string]interface{}, filter *Filter) error
	// Delete row(s) in the table with optional filter
	Delete(tableName string, filter *Filter) error

	Restart() error
	Close() error
}

// Entry pointer main
func GetGotdata() Gotdata {
	return mssql.GetInstance()
}

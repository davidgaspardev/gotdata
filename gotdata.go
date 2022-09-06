package gotdata

import "gotdata/helpers"

type Filter = helpers.Filter
type Where = helpers.Where

type Gotdata interface {
	// Execute T-SQL statement
	Exec(tSql string) error
	// Query data with T-SQL statement
	Query(tSql string) ([]map[string]interface{}, error)
	// Insert data in the table
	Write(tableName string, data map[string]interface{}) error
	// Select columns (attributes) from table
	Read(tableName string, attributes []string) ([]map[string]interface{}, error)
	// Select columns (attributes) from table with where statement
	ReadWithFilter(tableName string, attributes []string, filter *Filter) ([]map[string]interface{}, error)
	// Update row(s) in the table with where statement
	Update(tableName string, data map[string]interface{}, filter *Filter) error
	// Delete row(s) in the table with where statement
	Delete(tableName string, filter *Filter) error
}

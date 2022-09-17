package mssql

import (
	"fmt"
	"gotdata/helpers"
	"gotdata/tools"

	"github.com/davidgaspardev/golog"
)

func (io *_DatabaseMssql) SetLogger(showLog bool) {
	io.logger = showLog
}

func (io *_DatabaseMssql) Query(query string) ([]map[string]interface{}, error) {
	data, err := io.read(nil, query)
	if err != nil && io.logger {
		golog.Error("Gotdata", err)
	}

	return data, err
}

func (io *_DatabaseMssql) Exec(statement string) error {
	_, err := io.db.Exec(statement)
	if err != nil && io.logger {
		golog.Error("Gotdata", err)
	}

	return err
}

func (io *_DatabaseMssql) Count(tableName string) (uint, error) {
	var tSqlBuilder tools.TSqlBuilder
	tSqlBuilder.SelectCountAll()
	tSqlBuilder.From(tableName)
	tSqlStatement := tSqlBuilder.Done()

	rows, err := io.db.Query(tSqlStatement)
	if err != nil {
		if io.logger {
			golog.Error("Gotdata", err)
		}
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		err = fmt.Errorf("could not count table (%s)", tSqlStatement)
		if io.logger {
			golog.Error("Gotdata", err)
		}
		return 0, err
	}

	var count uint
	err = rows.Scan(&count)
	if err != nil && io.logger {
		golog.Error("Gotdata", err)
	}

	return count, err
}

func (io *_DatabaseMssql) Write(tableName string, data map[string]interface{}) error {
	var tSqlBuilder tools.TSqlBuilder
	tSqlBuilder.InsertInto(tableName, data)
	tSqlStatement := tSqlBuilder.Done()

	_, err := io.db.Exec(tSqlStatement)
	if err != nil && io.logger {
		golog.Error("Gotdata", err)
	}

	return err
}

func (io *_DatabaseMssql) Read(tableName string, attributes []string) ([]map[string]interface{}, error) {
	var tSqlBuilder tools.TSqlBuilder
	tSqlBuilder.SelectColumns(attributes)
	tSqlBuilder.From(tableName)
	tSqlStatement := tSqlBuilder.Done()

	data, err := io.read(attributes, tSqlStatement)
	if err != nil && io.logger {
		golog.Error("Gotdata", err)
	}

	return data, err
}

func (io *_DatabaseMssql) ReadWithFilter(tableName string, attributes []string, filter *helpers.Filter) ([]map[string]interface{}, error) {
	var tSqlBuilder tools.TSqlBuilder
	tSqlBuilder.SelectColumns(attributes)
	tSqlBuilder.From(tableName)
	if filter != nil {
		io.buildFilter(&tSqlBuilder, filter)
	}
	tSqlStatement := tSqlBuilder.Done()

	data, err := io.read(attributes, tSqlStatement)
	if err != nil && io.logger {
		golog.Error("Gotdata", err)
	}

	return data, err
}

func (io *_DatabaseMssql) Update(tableName string, data map[string]interface{}, filter *helpers.Filter) error {
	var tSqlBuilder tools.TSqlBuilder
	tSqlBuilder.Update(tableName, data)
	if filter != nil {
		io.buildFilter(&tSqlBuilder, filter)
	}
	tSqlStatement := tSqlBuilder.Done()

	_, err := io.db.Exec(tSqlStatement)
	if err != nil && io.logger {
		golog.Error("Gotdata", err)
	}

	return err
}

func (io *_DatabaseMssql) Delete(tableName string, filter *helpers.Filter) error {
	var tSqlBuilder tools.TSqlBuilder
	tSqlBuilder.Delete(tableName)
	if filter != nil {
		io.buildFilter(&tSqlBuilder, filter)
	}
	tSqlStatement := tSqlBuilder.Done()

	_, err := io.db.Exec(tSqlStatement)
	if err != nil && io.logger {
		golog.Error("Gotdata", err)
	}

	return err
}

func (io *_DatabaseMssql) read(columns []string, query string) ([]map[string]interface{}, error) {

	numberOfRows, err := io.countSubSelect(query)
	if err != nil {
		return nil, err
	}

	// Get the data (rows) from database
	rows, err := io.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if len(columns) == 0 {
		columns, err = rows.Columns()
		if err != nil {
			return nil, err
		}
	}

	// Allocate RAM space to store these rows of data
	allData := make([]map[string]interface{}, numberOfRows)

	columnsLen := uint(len(columns))
	var index uint
	for rows.Next() {
		columnsTmp := make([]interface{}, len(columns))
		columnsPointers := make([]interface{}, len(columns))

		for i := uint(0); i < columnsLen; i++ {
			columnsPointers[i] = &columnsTmp[i]
		}

		if err = rows.Scan(columnsPointers...); err != nil {
			return nil, err
		}

		data := make(map[string]interface{})
		for index, key := range columns {
			if len(key) > 7 && key[:7] == "_IGNORE" {
				continue
			}

			tmp := columnsPointers[index].(*interface{})
			data[key] = *tmp
		}

		allData[index] = data
		index++
	}

	// Return allocated data
	return allData, nil
}

// Count rows from subselect
func (io *_DatabaseMssql) countSubSelect(subSelect string) (uint, error) {
	var tSqlBuilder tools.TSqlBuilder

	tSqlBuilder.SelectCountAll()
	tSqlBuilder.FromSubSelect(subSelect[:len(subSelect)-1])
	tSqlStmt := tSqlBuilder.Done()

	rows, err := io.db.Query(tSqlStmt)
	if err != nil {
		err = fmt.Errorf("%s (%s)", err.Error(), tSqlStmt)
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		err = fmt.Errorf("could not count subSelect (%s)", tSqlStmt)
		return 0, err
	}

	var count uint
	err = rows.Scan(&count)

	return count, err
}

func (io *_DatabaseMssql) buildFilter(tSqlBuilder *tools.TSqlBuilder, filter *helpers.Filter) {
	numberOfWheres := len(filter.Wheres)

	if numberOfWheres > 0 {
		tSqlBuilder.Where()

		for i := 0; i < numberOfWheres; i++ {
			where := filter.Wheres[i]

			if i > 0 {
				tSqlBuilder.And()
			}

			switch where.Operator {
			case "=":
				tSqlBuilder.Equal(where.Attribute, where.Value)
			case "!=", "<>":
				tSqlBuilder.IsNotEqual(where.Attribute, where.Value)
			case ">":
				tSqlBuilder.GreaterThan(where.Attribute, where.Value)
			case ">=":
				tSqlBuilder.GreaterThanOrEqualTo(where.Attribute, where.Value)
			case "<":
				tSqlBuilder.LessThan(where.Attribute, where.Value)
			case "<=":
				tSqlBuilder.LessThanOrEqualTo(where.Attribute, where.Value)
			case "LIKE":
				tSqlBuilder.Like(where.Attribute, where.Value)
			case "IS NULL":
				tSqlBuilder.IsNull(where.Attribute)
			case "IS NOT NULL":
				tSqlBuilder.IsNotNull(where.Attribute)
			}
		}
	}

	hasOrder := len(filter.Orders) > 0

	if filter.Page > 0 {
		if hasOrder {
			hasOrder = false
			tSqlBuilder.OrderByColumns(filter.Orders)
		} else {
			tSqlBuilder.OrderBy("(SELECT NULL)")
		}

		tSqlBuilder.OffSet((filter.Page - 1) * /* max rows by page */ 50).FetchNext( /* max rows by page */ 50)
	}

	if hasOrder {
		tSqlBuilder.OrderByColumns(filter.Orders)
	}
}

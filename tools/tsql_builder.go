package tools

import (
	"fmt"
	"strings"
)

// Transact-SQL Builder
// See: https://docs.microsoft.com/en-us/sql/t-sql/language-reference?view=sql-server-ver16
type TSqlBuilder struct {
	statement string
}

// Add statement
func (tSqlBuilder *TSqlBuilder) addStatement(stmt string) {
	if tSqlBuilder.statement == "" {
		tSqlBuilder.statement = stmt
	} else {
		tSqlBuilder.statement = fmt.Sprint(tSqlBuilder.statement, " ", stmt)
	}
}

// ---------------------- General ----------------------

func (tSqlBuidler *TSqlBuilder) SubQuery(inside func(*TSqlBuilder)) *TSqlBuilder {
	tSqlBuidler.addStatement("(")
	inside(tSqlBuidler)
	tSqlBuidler.addStatement(")")
	return tSqlBuidler
}

// ---------------------- SELECT Clouse (Transact-SQL) ----------------------
//
// See: https://docs.microsoft.com/en-us/sql/t-sql/queries/select-clause-transact-sql?view=sql-server-ver16

func (tSqlBuilder *TSqlBuilder) Select(selectList []string) *TSqlBuilder {
	if len(selectList) > 0 {
		statement := fmt.Sprintf("SELECT %s", strings.Join(selectList, ","))
		tSqlBuilder.addStatement(statement)
	} else {
		return tSqlBuilder.SelectAll()
	}
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) SelectAll() *TSqlBuilder {
	tSqlBuilder.statement = "SELECT ALL"
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) SelectColumns(columns []string) *TSqlBuilder {
	columnsFormatted := make([]string, len(columns))
	if len(columns) > 0 {
		for i := 0; i < len(columns); i++ {
			columnsFormatted[i] = fmt.Sprintf("[%s]", columns[i])
		}
		statement := fmt.Sprintf("SELECT %s", strings.Join(columnsFormatted, ","))
		tSqlBuilder.addStatement(statement)
	} else {
		return tSqlBuilder.SelectAll()
	}
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) SelectDistinct(selectList []string) *TSqlBuilder {
	if len(selectList) > 0 {
		statement := fmt.Sprintf("SELECT DISTINCT %s", strings.Join(selectList, ","))
		tSqlBuilder.addStatement(statement)
	} else {
		return tSqlBuilder.SelectDistinct([]string{"*"})
	}
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) SelectDistinctColumn(columns []string) *TSqlBuilder {
	columnsFormatted := make([]string, len(columns))
	if len(columns) > 0 {
		for i := 0; i < len(columns); i++ {
			columnsFormatted[i] = fmt.Sprintf("[%s]", columns[i])
		}
		statement := fmt.Sprintf("SELECT DISTINCT %s", strings.Join(columnsFormatted, ","))
		tSqlBuilder.addStatement(statement)
	} else {
		return tSqlBuilder.SelectDistinct([]string{"*"})
	}
	return tSqlBuilder
}

// Count rows from table
func (tSqlBuilder *TSqlBuilder) SelectCountAll() *TSqlBuilder {
	statement := "SELECT COUNT(*)"

	tSqlBuilder.statement = statement

	return tSqlBuilder
}

// ---------------------- FROM Clouse (Transact-SQL) ----------------------
//
// See: https://docs.microsoft.com/en-us/sql/t-sql/queries/from-transact-sql?view=sql-server-ver16

func (tSqlBuilder *TSqlBuilder) From(tableSource string) *TSqlBuilder {
	statement := "FROM [" + tableSource + "]"
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

func (tSqlBuidler *TSqlBuilder) FromSubSelect(subSelect string) *TSqlBuilder {
	statement := "FROM (" + subSelect + ") AS TAB"
	tSqlBuidler.addStatement(statement)
	return tSqlBuidler
}

func (tSqlBuilder *TSqlBuilder) WithNolock() *TSqlBuilder {
	statement := "WITH (NOLOCK)"
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

// ---------------------- WHERE Clouse (Transact-SQL) ----------------------
//
// See: https://docs.microsoft.com/en-us/sql/t-sql/queries/where-transact-sql?view=sql-server-ver16

func (tSqlBuilder *TSqlBuilder) Where() *TSqlBuilder {
	tSqlBuilder.addStatement("WHERE")
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) And() *TSqlBuilder {
	tSqlBuilder.addStatement("AND")
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) Equal(column string, value interface{}) *TSqlBuilder {
	statement := fmt.Sprintf("[%s] = '%v'", column, value)
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) IsNotEqual(column string, value interface{}) *TSqlBuilder {
	statement := fmt.Sprintf("[%s] <> '%v'", column, value)
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) GreaterThan(column string, value interface{}) *TSqlBuilder {
	statement := fmt.Sprintf("[%s] > '%v'", column, value)
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) GreaterThanOrEqualTo(column string, value interface{}) *TSqlBuilder {
	statement := fmt.Sprintf("[%s] >= '%v'", column, value)
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) LessThan(column string, value interface{}) *TSqlBuilder {
	statement := fmt.Sprintf("[%s] < '%v'", column, value)
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) LessThanOrEqualTo(column string, value interface{}) *TSqlBuilder {
	statement := fmt.Sprintf("[%s] <= '%v'", column, value)
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) Between(column string, startValue interface{}, endValue interface{}) *TSqlBuilder {
	statement := fmt.Sprintf("[%s] BETWEEN '%v' AND '%v'", column, startValue, endValue)
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) Like(column string, value interface{}) *TSqlBuilder {
	statement := fmt.Sprintf("[%s] LIKE '%%%v%%'", column, value)
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) IsNull(column string) *TSqlBuilder {
	statement := fmt.Sprintf("[%s] IS NULL", column)
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

func (tSqlBuilder *TSqlBuilder) IsNotNull(column string) *TSqlBuilder {
	statement := fmt.Sprintf("[%s] IS NOT NULL", column)
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

// ---------------------- ORDER BY Clouse (Transact-SQL) ----------------------
//
// See: https://docs.microsoft.com/en-us/sql/t-sql/queries/select-order-by-clause-transact-sql?view=sql-server-ver16

func (tSqlBuidler *TSqlBuilder) OrderBy(order string) *TSqlBuilder {
	statement := "ORDER BY " + order
	tSqlBuidler.addStatement(statement)
	return tSqlBuidler
}

func (tSqlBuidler *TSqlBuilder) OrderByColumn(column string) *TSqlBuilder {
	statement := "ORDER BY [" + column + "]"
	tSqlBuidler.addStatement(statement)
	return tSqlBuidler
}

func (tSqlbuilder *TSqlBuilder) OrderByColumns(columns []string) *TSqlBuilder {
	for i := 0; i < len(columns); i++ {
		columns[i] = "[" + columns[i] + "]"
	}
	statement := "ORDER BY " + strings.Join(columns, ", ")
	tSqlbuilder.addStatement(statement)
	return tSqlbuilder
}

func (tSqlBuilder *TSqlBuilder) OffSet(rows uint) *TSqlBuilder {
	statement := fmt.Sprintf("OFFSET %d ROWS", rows)
	tSqlBuilder.addStatement(statement)
	return tSqlBuilder
}

func (tSqlBuidler *TSqlBuilder) FetchNext(rows uint) *TSqlBuilder {
	statement := fmt.Sprintf("FETCH NEXT %d ROWS ONLY", rows)
	tSqlBuidler.addStatement(statement)
	return tSqlBuidler
}

// ---------------------- INSERT Clouse (Transact-SQL) ----------------------
//
// See: https://docs.microsoft.com/en-us/sql/t-sql/statements/insert-sql-graph?view=sql-server-ver16

func (tSqlBuilder *TSqlBuilder) InsertInto(tableName string, data map[string]interface{}) *TSqlBuilder {
	statement := fmt.Sprintf("INSERT INTO [%s]", tableName)
	columns := make([]string, len(data))
	values := make([]string, len(data))

	var index uint8
	for column, value := range data {
		columns[index] = fmt.Sprintf("[%s]", column)
		values[index] = fmt.Sprintf("'%v'", value)
		index++
	}

	statement = fmt.Sprintf("%s (%s) VALUES (%s)", statement, strings.Join(columns, ","), strings.Join(values, ","))

	tSqlBuilder.addStatement(statement)

	return tSqlBuilder
}

// ---------------------- UPDATE Clouse (Transact-SQL) ----------------------
//
// See: https://docs.microsoft.com/en-us/sql/t-sql/queries/update-transact-sql?view=sql-server-ver16

func (tSqlBuilder *TSqlBuilder) Update(tableName string, data map[string]interface{}) *TSqlBuilder {
	statement := "UPDATE [" + tableName + "] SET "

	var index uint8
	for key, value := range data {
		if index > 0 {
			statement += " , "
		}
		statement += fmt.Sprintf("[%s] = '%v'", key, value)

		index++
	}

	tSqlBuilder.statement = statement
	return tSqlBuilder
}

// ---------------------- DELETE Clouse (Transact-SQL) ----------------------
//
// See: https://docs.microsoft.com/en-us/sql/t-sql/statements/delete-transact-sql?view=sql-server-ver16

func (tSqlBuidler *TSqlBuilder) Delete(tableName string) *TSqlBuilder {
	statement := "DELETE [" + tableName + "]"
	tSqlBuidler.addStatement(statement)
	return tSqlBuidler
}

// ---------------------- QUERY DONE ---------------------- //

func (tSqlBuilder *TSqlBuilder) Done() string {
	return fmt.Sprint(tSqlBuilder.statement, ";")
}

func (tSqlBuilder *TSqlBuilder) DoneTransaction() string {
	return fmt.Sprintf("BEGIN TRANSACTION; %s; COMMIT;", tSqlBuilder.statement)
}

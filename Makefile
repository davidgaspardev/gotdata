PACKAGE_NAME = github.com/davidgaspardev/gotdata

DIR_EXAMPLES = examples
DIR_MSSQL = $(PACKAGE_NAME)/mssql
DIR_TOOLS = $(PACKAGE_NAME)/tools

run_examples:
	go run $(DIR_EXAMPLES)/main.go

test: test_tools test_mssql

test_tools:
	go test $(DIR_TOOLS)

test_tools_verbose:
	go test -v $(DIR_TOOLS)

test_mssql:
	go test -p 1 $(DIR_MSSQL)

test_mssql_verbose:
	go test -p 1 -v $(DIR_MSSQL)
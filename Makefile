PACKAGE_NAME = gotdata

DIR_EXAMPLES = examples
DIR_MSSQL = $(PACKAGE_NAME)/mssql
DIR_TOOLS = $(PACKAGE_NAME)/tools

run_examples:
	go run $(DIR_EXAMPLES)/main.go

test:
	go test $(DIR_TOOLS)
	go test $(DIR_MSSQL)

test_tools:
	go test $(DIR_TOOLS)

test_tools_verbose:
	go test -v $(DIR_TOOLS)

test_mssql:
	go test $(DIR_MSSQL)

test_mssql_verbose:
	go test -v $(DIR_MSSQL)
PACKAGE_NAME = gotdata

DIR_EXAMPLES = examples
DIR_TOOLS = $(PACKAGE_NAME)/tools

run_examples:
	go run $(DIR_EXAMPLES)/main.go

tests:
	go test $(DIR_TOOLS)
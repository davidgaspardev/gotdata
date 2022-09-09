PACKAGE_NAME = gotdata

DIR_EXAMPLES = $(PACKAGE_NAME)/examples
DIR_TOOLS = $(PACKAGE_NAME)/tools

examples:
	go run $(DIR_EXAMPLES)/main.go

tests:
	go test $(DIR_TOOLS)
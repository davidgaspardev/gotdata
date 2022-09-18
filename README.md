# Gotdata

Gotdata is a minimal connection with SQL Server database, using T-SQL to read and write data.

## Install

Command to include the package in your project:

```bash
go get github.com/davidgaspardev/gotdata
```

Obs: run command in project root


## Environment

For authentication with Microsoft SQL Server, the following environment variables are required:

- `MSSQL_USERNAME`
- `MSSQL_PASSWORD`
- `MSSQL_HOST`
- `MSSQL_PORT`
- `MSSQL_INSTANCE`
- `MSSQL_DATABASE`

## Examples

Run examples code:

```bash
make run_examples
```

## Tests

Make all tests:

```bash
make tests
```
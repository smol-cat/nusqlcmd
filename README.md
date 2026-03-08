# Nusqlcmd WIP

## Description

Very minimal command line tool written in Go used to execute queries to SQL
server and outputs resulting data to stdout in JSON format. Currently supports
only SQL Server. Inspired by
[go-sqlcmd](https://github.com/microsoft/go-sqlcmd).

## Usage 

Usage:
  nusqlcmd \[OPTIONS\]

Application Options:
  -c, --config=            A path for configuration
  -s, --connection-string= Connection string to use to connect to database
  -p, --profile=           A profile to use to connect to target database
  -q, --query=             A query to execute

Help Options:
  -h, --help               Show this help message

## Application config

Stored in XDG config home directory: `~/.config/nusqlcmd/config.yaml`

Example:

```yaml
profiles:
  - name: coolProfile
    connectionString: sqlserver://username:supersecurepassword123!@localhost:1433?database=master
```

package main

import (
	"fmt"
	"os"

	_ "github.com/microsoft/go-mssqldb"
	"github.com/smol-cat/nusqlcmd/src/common"
	"github.com/smol-cat/nusqlcmd/src/core"
	"github.com/smol-cat/nusqlcmd/src/serialization"
)

type RowScanner struct {
}

func (r RowScanner) Scan(src any) error {
	return nil
}

func main() {
	dbConnection, err := core.ConnectToDb("sqlserver://sa:password123!@localhost:1433?database=master")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if dbConnection == nil {
		fmt.Println("Failed to connect to the DB")
		return
	}

	fmt.Println(os.Args[1])
	rows, err := dbConnection.Query(os.Args[1])

	var result = serialization.SerializeRowsToTable(rows)
	fmt.Println(result)

	common.PanicOnErr(err)
}

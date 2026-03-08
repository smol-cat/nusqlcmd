package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	_ "github.com/microsoft/go-mssqldb"
	"github.com/smol-cat/nusqlcmd/src/common"
	"github.com/smol-cat/nusqlcmd/src/config"
	"github.com/smol-cat/nusqlcmd/src/core"
	"github.com/smol-cat/nusqlcmd/src/serialization"
)

type RowScanner struct {
}

func (r RowScanner) Scan(src any) error {
	return nil
}

func main() {
	cmdParams, err := config.ReadFlags()
	if err != nil {
		if !flags.WroteHelp(err) {
			fmt.Println(err.Error())
		}

		os.Exit(1)
	}

	configPath := config.GetDefaultConfigPath()
	config, err := config.ReadConfig(configPath)
	common.PanicOnErr(err)

	dbConnection, err := core.ConnectToDb(config.Profiles[0].ConnectionString)
	common.PanicOnErr(err)

	if dbConnection == nil {
		fmt.Println("Failed to connect to the DB")
		return
	}

	rows, err := dbConnection.Query(cmdParams.Query)
	common.PanicOnErr(err)

	var result = serialization.SerializeToJson(rows)
	fmt.Print(result)

	common.PanicOnErr(err)
}

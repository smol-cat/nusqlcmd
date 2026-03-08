package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
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
	common.ExitOnErrFunc(err, 1, func(err error) {
		if !flags.WroteHelp(err) {
			fmt.Println(err.Error())
		}
	})

	configPath := config.GetDefaultConfigPath()
	config, err := config.ReadConfig(configPath)
	common.ExitOnErr(err, 1)

	dbConnection, err := core.ConnectToDb(config.Profiles[0].ConnectionString)
	common.ExitOnErr(err, 1)

	if dbConnection == nil {
		fmt.Println("Failed to connect to the DB")
		return
	}

	rows, err := dbConnection.Query(cmdParams.Query)
	common.ExitOnErr(err, 1)

	var result = serialization.SerializeToJson(rows)
	fmt.Print(result)

	common.ExitOnErr(err, 1)
}

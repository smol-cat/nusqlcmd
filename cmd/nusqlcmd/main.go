package main

import (
	"fmt"

	"github.com/jessevdk/go-flags"
	"github.com/smol-cat/nusqlcmd/internal/common"
	"github.com/smol-cat/nusqlcmd/internal/config"
	"github.com/smol-cat/nusqlcmd/internal/core"
	"github.com/smol-cat/nusqlcmd/internal/serialization"
)

func main() {
	configPath := config.GetDefaultConfigPath()
	appConfig, err := config.ReadAppConfig(configPath)
	common.ExitOnErr(err, 1)

	cmdParams, err := config.ReadFlags()
	common.ExitOnErrFunc(err, 1, func(err error) {
		if !flags.WroteHelp(err) {
			fmt.Println(err.Error())
		}
	})

	runtimeConfig, err := config.ConsolidateIntoRuntimeConfig(appConfig, cmdParams)
	common.ExitOnErr(err, 1)

	dbConnection, err := core.ConnectToDb(runtimeConfig.ConnectionString)
	common.ExitOnErr(err, 1)

	if dbConnection == nil {
		fmt.Println("Failed to connect to the DB")
		return
	}

	rows, err := dbConnection.Query(runtimeConfig.Query)
	common.ExitOnErr(err, 1)

	var result = serialization.SerializeToJson(rows)
	fmt.Print(result)

	common.ExitOnErr(err, 1)
}

package main

import (
	"errors"
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
		common.ExitOnErr(errors.New("Failed to connect to the DB"), 1)
	}

	rows, err := dbConnection.Query(runtimeConfig.Query)
	common.ExitOnErr(err, 1)

	result, err := serialization.SerializeToJson(rows)
	common.ExitOnErr(err, 1)

	fmt.Print(result)
}

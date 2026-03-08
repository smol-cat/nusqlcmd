package config

import (
	"errors"
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/goccy/go-yaml"
	"github.com/jessevdk/go-flags"
)

func GetDefaultConfigPath() string {
	return path.Join(xdg.ConfigHome, "nusqlcmd/config.yaml")
}

func ReadAppConfig(path string) (AppConfig, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return AppConfig{}, errors.New("Config file open error: " + err.Error())
	}

	config := AppConfig{}
	err = yaml.Unmarshal(bytes, &config)

	if err != nil {
		return AppConfig{}, errors.New("Config parse error: " + err.Error())
	}

	return config, nil
}

func ReadFlags() (CommandLineArgs, error) {
	cla := CommandLineArgs{}
	_, err := flags.ParseArgs(&cla, os.Args[1:])
	if err != nil {
		return cla, err
	}

	if cla.ConfigPath == "" {
		cla.ConfigPath = GetDefaultConfigPath()
	}

	if cla.ConnectionString == "" && cla.Profile == "" {
		err = errors.New("Need to provide either connection string (-cs) or a profile (-p)")
	}

	return cla, err
}

func getConnStringFromConfig(config AppConfig, profile string) (string, error) {
	if len(config.Profiles) == 0 {
		return "", errors.New("There are no profiles defined in the configuration file")
	}

	for i := range config.Profiles {
		if config.Profiles[i].Name == profile {
			return config.Profiles[i].ConnectionString, nil
		}
	}

	return "", errors.New("Profile with name '" + profile + "' was not found in configuration")
}

func ConsolidateIntoRuntimeConfig(config AppConfig, cla CommandLineArgs) (RuntimeConfig, error) {
	runtimeConf := RuntimeConfig{}
	runtimeConf.Query = cla.Query
	runtimeConf.ConnectionString = cla.ConnectionString

	if runtimeConf.ConnectionString == "" {
		connectionString, err := getConnStringFromConfig(config, cla.Profile)
		if err != nil {
			return runtimeConf, err
		}

		runtimeConf.ConnectionString = connectionString
	}

	return runtimeConf, nil
}

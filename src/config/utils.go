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

func ReadConfig(path string) (Config, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return Config{}, errors.New("Config file open error: " + err.Error())
	}

	config := Config{}
	err = yaml.Unmarshal(bytes, &config)

	if err != nil {
		return Config{}, errors.New("Config parse error: " + err.Error())
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

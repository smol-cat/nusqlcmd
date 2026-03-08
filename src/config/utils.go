package config

import (
	"errors"
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/goccy/go-yaml"
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

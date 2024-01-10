package configuration

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ConStringMsDb       ConStringMS `yaml:"mssql"`
	AdAddress           string      `yaml:"ad_address"`
	GetACSUser          string      `yaml:"get_acsuser"`
	UpdateACSUser       string      `yaml:"update_acsuser"`
	GetPermissions      string      `yaml:"get_permissions"`
	GetGuestPermissions string      `yaml:"get_guest_permissions"`
	ServiceId           *int        `yaml:"service_id"`
	ServerAddress       string      `yaml:"server_address"`
	UrlPrefix           string      `yaml:"url_prefix"`
	GinMode             string      `yaml:"gin_mode"`
}

type ConStringMS struct {
	Server   string `yaml:"server"`
	UserID   string `yaml:"user_id"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

var (
	GlobalConfig Config = initConfig()
)

func initConfig() Config {
	configFiles := []string{"configuration/config.yml", "am-fuel-gas-webapi.conf.yml"}
	var configName string
	var config Config

	for _, configFile := range configFiles {
		if _, err := os.Stat(configFile); err == nil {
			configName = configFile
			break
		}
	}

	data, err := os.ReadFile(configName)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

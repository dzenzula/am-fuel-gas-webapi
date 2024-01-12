package configuration

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ConStringMsDb ConStringMS `yaml:"mssql"`
	ConStringPgDb ConStringPG `yaml:"pg"`
	Permissions   Permissions `yaml:"permissions"`
	AdAddress     string      `yaml:"ad_address"`
	Querries      Querries    `yaml:"querries"`
	ServiceId     *int        `yaml:"service_id"`
	ServerAddress string      `yaml:"server_address"`
	GinMode       string      `yaml:"gin_mode"`
}

type Querries struct {
	GetACSUser          string `yaml:"get_acsuser"`
	UpdateACSUser       string `yaml:"update_acsuser"`
	GetPermissions      string `yaml:"get_permissions"`
	GetGuestPermissions string `yaml:"get_guest_permissions"`
	InsertParametrs     string `yaml:"insert_parametrs"`
	GetMeasuringsIds    string `yaml:"get_measuringids"`
	GetData             string `yaml:"get_data"`
}

type ConStringMS struct {
	Server   string `yaml:"server"`
	UserID   string `yaml:"user_id"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type ConStringPG struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type Permissions struct {
	Edit string `yaml:"edit"`
	Show string `yaml:"show"`
}

var (
	GlobalConfig Config = initConfig()
)

func initConfig() Config {
	executable, err := os.Executable()
	if err != nil {
		log.Fatalf(fmt.Sprintf("Unable to get executable path: %s", err))
	}
	configPath := filepath.Join(filepath.Dir(executable), "am-fuel-gas-webapi.conf.yml")
	configFiles := []string{"configuration/config.yml", configPath}
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

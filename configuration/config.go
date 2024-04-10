package configuration

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ConStringPgDb ConStringPG `yaml:"pg"`
	Permissions   Permissions `yaml:"permissions"`
	ServerAddress string      `yaml:"server_address"`
	GinMode       string      `yaml:"gin_mode"`
	LogLevel      string      `yaml:"log_level"`
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
	Edit       string `yaml:"edit"`
	Show       string `yaml:"show"`
	Calculate  string `yaml:"calculate"`
	EditScales string `yaml:"edit_scales"`
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

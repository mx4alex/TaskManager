package config

import (
	"github.com/spf13/viper"
)

const (
	ConfigFilePath = "./"
	ConfigFileName = "config"
)

type Config struct {
	StorageType string 			`mapstructure:"storage_type"`
	HttpPort  	string			`mapstructure:"http_port"`
	Interface   string  		`mapstructure:"interface"`
	Postgres 	PostgresConfig  `mapstructure:"postgres"`
}

type PostgresConfig struct {
	Host 	 string  `mapstructure:"host"`
	Port 	 int 	 `mapstructure:"port"`
	User 	 string	 `mapstructure:"user"`
	Password string  `mapstructure:"password"`
	DBName 	 string  `mapstructure:"dbname"`
	SSLMode	 string  `mapstructure:"sslmode"`
}

func New() (Config, error) {
	vpr := viper.New()
	vpr.AddConfigPath(ConfigFilePath)
	vpr.SetConfigName(ConfigFileName)

	if err := vpr.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var result Config
	if err := vpr.Unmarshal(&result); err != nil {
		return Config{}, err
	}

	return result, nil
}
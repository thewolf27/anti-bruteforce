package config

import (
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	AppConfig
	LoggerConfig
	ServerConfig
	StorageConfig
}

type AppConfig struct {
	NumberOfAttemptsForLogin    int
	NumberOfAttemptsForPassword int
	NumberOfAttemptsForIP       int
}

type LoggerConfig struct {
	Level string
	File  string
}

type ServerConfig struct {
	Address string
}

type StorageConfig struct {
	Dsn string
}

func NewConfig() *Config {
	parseFlags()

	configFolder := viper.GetString("configFolder")
	viper.SetConfigType("yml")
	viper.AddConfigPath(configFolder)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	var config Config
	viper.Unmarshal(&config)

	config.StorageConfig.Dsn = os.Getenv("DSN")

	return &config
}

func parseFlags() {
	pflag.String("configFolder", "./configs", "path to configs folder")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

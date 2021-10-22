package config

import (
	"log"

	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Jwt      JwtConfig      `yaml:"jwt"`
}

func GetConfig() *Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	// config ENV

	// Server
	configEnv("server.host", "SERVER_HOST")
	configEnv("server.port", "SERVER_PORT")

	// Database
	configEnv("database.host", "DB_HOST")
	configEnv("database.port", "DB_PORT")
	configEnv("database.user", "DB_USER")
	configEnv("database.password", "DB_PASSWORD")
	configEnv("database.dbname", "DB_NAME")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var configuration Configuration

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &configuration
}

func configEnv(input ...string) {
	err := viper.BindEnv(input...)
	if err != nil {
		log.Fatalf("Error BindEnv, %s", err)
	}
}

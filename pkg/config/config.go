package config

import (
	"github.com/spf13/viper"
	"log"
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

	var configuration Configuration

	// config ENV
	configEnv("server.port", "SERVER_PORT")
	configEnv("server.host", "SERVER_HOST")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

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

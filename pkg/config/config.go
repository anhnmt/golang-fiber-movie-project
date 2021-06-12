package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
}

// YamlConfig This maps the configuration in the yaml file
// into a struct
type YamlConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

func ReadYaml(path string) *YamlConfig {
	if path == "" {
		path = defaultYamlConfigPath()
	}

	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	defer func() { _ = f.Close() }()

	var cfg YamlConfig
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Printf("error reading yaml file into config struct: %s\n", err)
		os.Exit(2)
	}
	return &cfg
}

func defaultYamlConfigPath() string {
	// Reads the path of the current executable
	// goes up 2 directories and appends config.yaml
	// to the path.
	ex, err := os.Executable()
	if err != nil {
		log.Printf("error encountered reading path: %s\n", err)
		os.Exit(2)
	}

	filename := "config.yml"
	dir := path.Dir(path.Dir(ex))
	dir = path.Join(dir, filename)
	return dir
}

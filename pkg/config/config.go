package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

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

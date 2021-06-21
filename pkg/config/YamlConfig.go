package config

// YamlConfig This maps the configuration in the yaml file
// into a struct
type YamlConfig struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Jwt      JwtConfig      `yaml:"jwt"`
}

func GetConfig() *YamlConfig {
	return ReadYaml("")
}

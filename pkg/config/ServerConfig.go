package config

type ServerConfig struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Logger bool   `yaml:"logger"`
}

func GetServer() *ServerConfig {
	return &GetConfig().Server
}

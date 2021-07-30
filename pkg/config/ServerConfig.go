package config

type ServerConfig struct {
	Name    string `yaml:"name"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	Logger  bool   `yaml:"logger"`
	Prefork bool   `yaml:"prefork"`
}

func GetServer() *ServerConfig {
	return &GetConfig().Server
}

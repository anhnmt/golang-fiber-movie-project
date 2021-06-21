package config

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

func GetRedis() *RedisConfig {
	return &GetConfig().Redis
}

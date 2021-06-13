package config

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"dbname"`
	Migrate  bool   `yaml:"migrate"`
	Logger   bool   `yaml:"logger"`
}

func GetDatabase() *DatabaseConfig {
	return &GetConfig().Database
}

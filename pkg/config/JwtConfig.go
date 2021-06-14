package config

type JwtConfig struct {
	Secret string `yaml:"secret"`
}

func GetJwt() *JwtConfig {
	return &GetConfig().Jwt
}

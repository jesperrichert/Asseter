package env

import (
	"go.Asseter/internal/util"
)

type Config struct {
	AllowRegister string
}

func NewConfig() *Config {
	return &Config{
		AllowRegister: util.EnvToConfigValue("ALLOW_REGISTER", "true"),
	}
}

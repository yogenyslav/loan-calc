// Package config общая конфигурация сервиса
package config

import (
	srvconfig "loan/internal/server/config"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config хранит все конфиги отдельных частей сервиса.
type Config struct {
	Server srvconfig.Config `yaml:"server"`
}

// MustNew конструктор для Config, вызовет panic() в случае ошибки создания.
func MustNew(path string) *Config {
	cfg := &Config{}
	if err := cleanenv.ReadConfig(path, cfg); err != nil {
		panic(err)
	}
	return cfg
}

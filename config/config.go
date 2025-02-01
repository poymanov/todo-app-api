package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type DB struct {
	Host     string `env-required:"true" yaml:"host"`
	Port     string `env-required:"true" yaml:"port"`
	Name     string `env-required:"true" yaml:"name"`
	User     string `env-required:"true" yaml:"user"`
	Password string `env-required:"true" yaml:"password"`
}

type Auth struct {
	Secret string `env-required:"true" yaml:"secret"`
}

type Config struct {
	DB   DB   `yaml:"db"`
	Auth Auth `yaml:"auth"`
}

func (db *DB) DbConnectionAsString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		db.User, db.Password, db.Host, db.Port, db.Name)
}

func NewConfig() *Config {
	cfg := &Config{}
	if err := cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		panic(fmt.Errorf("ошибка загрузки конфигурации: %w", err))
	}

	return cfg
}

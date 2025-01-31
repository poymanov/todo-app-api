package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type DB struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	DB DB `yaml:"db"`
}

func (db *DB) DbConnectionAsString() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		db.User, db.Password, db.Host, db.Port, db.Name)
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		return nil, fmt.Errorf("ошибка загрузки конфигурации: %w", err)
	}

	return cfg, nil
}

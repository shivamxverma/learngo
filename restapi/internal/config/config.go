package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type DbConfig struct {
	Host     string `yaml:"host" env:"DB_HOST" env-required:"true"`
	Port     int    `yaml:"port" env:"DB_PORT" env-required:"true"`
	User     string `yaml:"user" env:"DB_USER" env-required:"true"`
	Password string `yaml:"password" env:"DB_PASSWORD" env-required:"true"`
	Name     string `yaml:"name" env:"DB_NAME" env-required:"true"`
	Sslmode  string `yaml:"sslmode" env:"DB_SSLMODE" env-default:"disable"`
}

type Config struct {
	Env string   `yaml:"env" env:"ENV" env-required:"true"`
	Db  DbConfig `yaml:"db"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg
}

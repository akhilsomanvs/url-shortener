package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Name           string `yaml:"name"`
	Path           string `yaml:"path"`
	Version        string `yaml:"version"`
	CollectionName string `yaml:"collection"`
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
	Database    `yaml:"database"`
}

func (cfg Config) GetHostAddress() string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}
func (cfg Config) GetStorageAddress() string {
	return cfg.Database.Path
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	log.Println("Config Path : " + configPath)
	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags
		log.Println("Config Path : " + configPath)
		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist %s", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("Cannot read config file : %s", err.Error())
	}

	return &cfg

}

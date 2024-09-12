package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Database struct{
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	ServerPort int
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	port, _ := strconv.Atoi(os.Getenv("ServerPort"))

	return &Config{
		Database: struct{Host string; Port string; User string; Password string; Name string}{
			Host: os.Getenv("DBHost"),
			Port: os.Getenv("DBPort"),
			User: os.Getenv("DBUser"),
			Password: os.Getenv("DBPassword"),
			Name: os.Getenv("DBName"),
		},

		ServerPort: port,
	}
}

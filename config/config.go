package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database   Database `envconfig:"db" required:"true"`
	ServerPort int      `envconfig:"serverport" required:"true" default:"8080"`
}

type Database struct {
  Host     string `envconfig:"host" required:"true"`
  Port     string `envconfig:"port" required:"true"`
  User     string `envconfig:"user" required:"true"`
  Password string `envconfig:"password" required:"true"`
  Name     string `envconfig:"name" required:"true"`
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}

	// cfg := &Config{
	// 	Database: Database{
	// 		Host: "postgres",
	// 		Port: "5432",
	// 		User: "postgres",
	// 		Password: "098098",
	// 		Name: "card-project",
	// 	},
	// 	ServerPort: 8080,
	// }

	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		log.Fatal("Failed load envconfig " + err.Error())
	}

	return cfg
}

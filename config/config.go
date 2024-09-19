package config

import (
	"log"

	"github.com/joho/godotenv"
	// "github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database   Database `envconfig:"DB" required:"true"`
	ServerPort int      `envconfig:"PORT_ServerPort" required:"true" default:"8080"`
}

type Database struct {
	Host     string `envconfig:"DB_Host" required:"true"`
	Port     string `envconfig:"DB_Port" required:"true"`
	User     string `envconfig:"DB_User" required:"true"`
	Password string `envconfig:"DB_Password" required:"true"`
	Name     string `envconfig:"DB_Name" required:"true"`
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	cfg := &Config{
		Database: Database{
			Host: "postgres",
			Port: "5432",
			User: "postgres",
			Password: "098098",
			Name: "card-project",
		},
		ServerPort: 8080,
	}

	// if err := envconfig.Process("", cfg); err != nil {
	// 	log.Fatal("Failed load envconfig " + err.Error())
	// }

	return cfg
}

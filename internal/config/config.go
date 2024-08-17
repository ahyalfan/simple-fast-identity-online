package config

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/lpernett/godotenv"
)

type Config struct {
	Server    Server
	Databases Databases
}

type Server struct {
	Host string
	Port string
}
type Databases struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func Get() *Config {
	// kita bisa melakukan pengambilan env dengan godotenv
	// di file .env secara default
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error when load .env", err.Error())
	}

	return &Config{
		Server: Server{
			// disini kita coba pakai os untuk mengabil env secara global
			// kalau godotenv biasanya untuk file .env
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Databases: Databases{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
	}

}

package configuration

import (
	"log"

	"github.com/joho/godotenv"
)

func InitializeEnv() {
	//loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Println("Error while opening .env file", err.Error())
	}
}

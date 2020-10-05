package configuration

import (
	"fmt"

	"github.com/joho/godotenv"
)

func InitializeEnv() {
	//loads values from .env into the system
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error while opening .env file", err.Error())
	}
}

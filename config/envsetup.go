// models/setup.go
package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

// SetupDB : initializing mysql database
func ENVSetup() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

}

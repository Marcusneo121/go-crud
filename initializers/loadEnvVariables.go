package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

//Start a function with capital letter if you want to use across other files
func LoadEnvVariables() {
err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}
}
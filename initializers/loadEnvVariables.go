package initializers

import (
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
)

func LoadEnvVariables() {
	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

package initializers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


func LoadEnvVaribles() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "local"
	}

	envFile := fmt.Sprintf(".env.%s", env)

	err := godotenv.Load(envFile)
	if err != nil {
		panic(fmt.Sprintf("Error loading %s: %v", envFile, err))
	}
}
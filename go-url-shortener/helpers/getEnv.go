package helpers

import (
	"os"

	"github.com/joho/godotenv"
)

// InitializeEnvVars loads env variables on config.env
func InitializeEnvVars() error {
	err := godotenv.Load("config.env")
	return err
}

// GetEnv returns a variable on env config file.
func GetEnv(env string) string {
	result := os.Getenv(env)
	return result
}

// GetEnvMany returns two variables on env config file.
func GetEnvMany(env1, env2 string) (firstResult, secondResult string) {
	firstResult = os.Getenv(env1)
	secondResult = os.Getenv(env2)
	return
}

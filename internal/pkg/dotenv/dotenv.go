package dotenv

import (
	johoDotenv "github.com/joho/godotenv"
)

func LoadEnvironment() error {
	return johoDotenv.Load(".env")
}

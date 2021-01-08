package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv .env読み込み
func LoadEnv() error {
	currentDir, _ := os.Getwd()
	err := godotenv.Load(fmt.Sprintf("%s/.env", currentDir))
	if err != nil {
		return err
	}
	return nil
}

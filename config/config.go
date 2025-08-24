package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file undefined")
	}
}

func GetHost() string {
	return os.Getenv("HOST")
}

func GetPort() string {
	return os.Getenv("PORT")
}

func GetSize() (int64, error) {
	maxSizeStr := os.Getenv("MAX_SIZE")
	maxSize, err := strconv.ParseInt(maxSizeStr, 10, 64)
	if err != nil {
		return 0, errors.New("error convertation size to int64")
	}
	return maxSize, nil
}

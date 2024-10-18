package util

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DBDriver            string
	DBSource            string
	ServerAddress       string
	TokenSymmetricKey   string
	AccessTokenDuration time.Duration
}

func LoadConfig() (config Config, err error) {
	config.DBSource = os.Getenv("DB_SOURCE")
	if config.DBSource == "" {
		log.Fatal("DB_SOURCE environment variable is not set")
	}

	config.DBDriver = os.Getenv("DB_DRIVER")
	if config.DBDriver == "" {
		log.Fatal("DB_DRIVER environment variable is not set")
	}

	config.ServerAddress = os.Getenv("SERVER_ADDRESS")
	if config.ServerAddress == "" {
		log.Fatal("SERVER_ADDRESS environment variable is not set")
	}

	config.TokenSymmetricKey = os.Getenv("TOKEN_SYMMETRIC_KEY")
	if config.TokenSymmetricKey == "" {
		log.Fatal("TOKEN_SYMMETRIC_KEY environment variable is not set")
	}

	accessTokenDurationStr := os.Getenv("ACCESS_TOKEN_DURATION")
	if accessTokenDurationStr == "" {
		log.Fatal("ACCESS_TOKEN_DURATION environment variable is not set")
	}

	accessTokenDuration, err := strconv.ParseInt(accessTokenDurationStr, 10, 64)
	if err != nil {
		log.Fatal("Invalid ACCESS_TOKEN_DURATION value")
	}

	config.AccessTokenDuration = time.Duration(accessTokenDuration) * time.Second

	return
}

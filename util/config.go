package util

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

func LoadConfig(isOS bool) (config Config, err error) {
	if isOS {
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

		config.AccessTokenDuration, err = time.ParseDuration(accessTokenDurationStr)
		if err != nil {
			log.Fatal("Invalid ACCESS_TOKEN_DURATION value")
		}

		refreshTokenDuration := os.Getenv("REFRESH_TOKEN_DURATION")
		if refreshTokenDuration == "" {
			log.Fatal("REFRESH_TOKEN_DURATION environment variable is not set")
		}
		config.RefreshTokenDuration, err = time.ParseDuration(refreshTokenDuration)
		if err != nil {
			log.Fatal("Invalid REFRESH_TOKEN_DURATION value")
		}

	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("app")
		viper.SetConfigType("env")

		viper.AutomaticEnv()
		err = viper.ReadInConfig()
		if err != nil {
			return
		}
		err = viper.Unmarshal(&config)

	}

	return
}

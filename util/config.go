package util

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	DatabaseURL         string        `mapstructure:"DATABASE_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	//viper.AddConfigPath(path)
	//viper.SetConfigType("env")

	viper.AutomaticEnv()
	//err = viper.ReadInConfig()
	viper.GetString("DB_DRIVER")
	viper.GetString("DB_SOURCE")
	viper.GetString("SERVER_ADDRESS")
	viper.GetString("TOKEN_SYMMETRIC_KEY")
	viper.GetTime("ACCESS_TOKEN_DURATION")
	viper.GetString("DATABASE_URL")
	//
	//if dbDriver == "" || dbSource == "" || serverAddress == "" || tokenSymmetricKey == "" || accessTokenDuration == "" || databaseURL == "" {
	//	log.Fatal("One or more required environment variables are not set")
	//}
	//
	//if err != nil {
	//	return
	//}
	//err = viper.Unmarshal(&config)

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal(err)
	}

	if config.DBDriver == "" || config.DBSource == "" || config.ServerAddress == "" || config.TokenSymmetricKey == "" || config.DatabaseURL == "" || config.AccessTokenDuration == 0 {
		log.Fatal("One or more required environment variables are not set")
	}
	return

}

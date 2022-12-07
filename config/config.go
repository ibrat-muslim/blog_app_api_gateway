package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HttpPort            string
	UserServiceHost     string
	UserServiceGrpcPort string
}

func Load(path string) Config {
	err := godotenv.Load(path + "/.env") // load .env file if it exists
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}

	conf := viper.New()
	conf.AutomaticEnv()

	cfg := Config{
		HttpPort:            conf.GetString("HTTP_PORT"),
		UserServiceHost:     conf.GetString("USER_SERVICE_HOST"),
		UserServiceGrpcPort: conf.GetString("USER_SERVICE_GRPC_PORT"),
	}

	return cfg
}

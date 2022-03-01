package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	JwtPublicKey   string
	JwtPrivateKey  string
	JwtKeyId       string
	AuthGrpcServer string
}

var lock = &sync.Mutex{}

var config *Config

func GetConfig() (*Config, error) {
	if config == nil {
		lock.Lock()
		defer lock.Unlock()
		if config == nil {
			conf := new(Config)

			err := godotenv.Load()

			if err != nil {
				log.Fatal("Error loading .env file")
			}

			setFromEnv(conf)

			config = conf
		}
	}

	return config, nil
}

func setFromEnv(cfg *Config) {
	cfg.AuthGrpcServer = os.Getenv("AUTH_GRPC_SERVER")
	cfg.JwtPublicKey = os.Getenv("JWT_PUBLIC_KEY")
	cfg.JwtPrivateKey = os.Getenv("JWT_PRIVATE_KEY")
	cfg.JwtKeyId = os.Getenv("JWT_KEY_ID")
}

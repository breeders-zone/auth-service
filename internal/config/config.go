package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	OauthBase         string
	ClientCallbackUrl string

	JwtPublicKey   string
	JwtPrivateKey  string
	JwtKeyId       string
	AuthGrpcServer string

	VkCleintId     string
	VkClientSercet string
}

var lock = &sync.Mutex{}

var config *Config

func GetConfig() (*Config, error) {
	if config == nil {
		lock.Lock()
		defer lock.Unlock()
		if config == nil {
			conf := new(Config)

			godotenv.Load()

		

			setFromEnv(conf)

			config = conf
		}
	}

	return config, nil
}

func setFromEnv(cfg *Config) {
	cfg.OauthBase = os.Getenv("OAUTH_BASE")
	cfg.ClientCallbackUrl = os.Getenv("CLIENT_CALLBACK_URL")

	cfg.AuthGrpcServer = os.Getenv("AUTH_GRPC_SERVER")
	cfg.JwtPublicKey = os.Getenv("JWT_PUBLIC_KEY")
	cfg.JwtPrivateKey = os.Getenv("JWT_PRIVATE_KEY")
	cfg.JwtKeyId = os.Getenv("JWT_KEY_ID")

	cfg.VkCleintId = os.Getenv("VKONTAKTE_CLIENT_ID")
	cfg.VkClientSercet = os.Getenv("VKONTAKTE_CLIENT_SECRET")
}

package jwk

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/breeders-zone/auth-service/internal/config"
	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
)

func GetKey() (jwk.Key, error) {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Not load config")
	}

	prvKey, err := ioutil.ReadFile(conf.JwtPrivateKey)
	if err != nil {
		log.Fatalln(err)
	}

	raw, err := jwt.ParseRSAPrivateKeyFromPEM(prvKey)
	if err != nil {
		fmt.Printf("failed to generate new RSA privatre key: %s\n", err)
		return nil, err;
	}

	key, err := jwk.New(raw)
	if err != nil {
		fmt.Printf("failed to create symmetric key: %s\n", err)
		return nil, err;
	}

	if _, ok := key.(jwk.RSAPrivateKey); !ok {
		fmt.Printf("expected jwk.SymmetricKey, got %T\n", key)
		return nil, err;
	}
	
	key.Set(jwk.KeyIDKey, conf.JwtKeyId)

	return key, nil
}
package jwt

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/breeders-zone/auth-service/internal/config"
	"github.com/golang-jwt/jwt"
)

func Create(ttl time.Duration, content interface{}) (string, error) {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Not load config")
	}

	prvKey, err := ioutil.ReadFile(conf.JwtPrivateKey)
	if err != nil {
		log.Fatalln(err)
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(prvKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["uid"] = content             // Our custom data.
	claims["exp"] = now.Add(ttl).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()          // The time at which the token was issued.
	claims["nbf"] = now.Unix()          // The time before which the token must be disregarded.

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = "foo"

	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return tokenStr, nil
}

func Validate(token string) (interface{}, error) {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Not load config")
	}

	pubKey, err := ioutil.ReadFile(conf.JwtPublicKey)
	if err != nil {
		log.Fatalln(err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		return "", fmt.Errorf("validate: parse key: %w", err)
	}
 
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
 
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}
 
	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("validate: invalid")
	}
 
	return claims["uid"], nil
}

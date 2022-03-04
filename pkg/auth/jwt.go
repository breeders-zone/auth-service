package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenManager struct {
	JwkId      string
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

func NewTokenManager(jwkKey string, pubKeyFile string, prvKeyFile string) (*TokenManager, error) {

	if prvKeyFile == "" && prvKeyFile == "" {
		key, err := rsa.GenerateKey(rand.Reader, 4096)

		if err != nil {
			return nil, err
		}

		pubKey := &key.PublicKey
		prvKey := key

		return &TokenManager{
			jwkKey,
			pubKey,
			prvKey,
		}, nil
	}

	prvKeyBytes, err := ioutil.ReadFile(prvKeyFile)
	if err != nil {
		log.Fatalln(err)
	}

	prvKey, err := jwt.ParseRSAPrivateKeyFromPEM(prvKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("create: parse key: %w", err)
	}

	pubKeyBytes, err := ioutil.ReadFile(pubKeyFile)
	if err != nil {
		log.Fatalln(err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	return &TokenManager{
		jwkKey,
		pubKey,
		prvKey,
	}, nil

}

func (t *TokenManager) Create(ttl time.Duration, id int32, roles interface{}) (string, error) {
	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["uid"] = id    
	claims["roles"] = roles        
	claims["exp"] = now.Add(ttl).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()          // The time at which the token was issued.
	claims["nbf"] = now.Unix()          // The time before which the token must be disregarded.

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = "foo"

	tokenStr, err := token.SignedString(t.PrivateKey)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return tokenStr, nil
}

func (t *TokenManager) Validate(token string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return t.PrivateKey, nil
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

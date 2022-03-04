package auth

import (
	"fmt"
	"github.com/lestrrat-go/jwx/jwk"
)

func (t *TokenManager) GetJwk() (jwk.Key, error) {
	key, err := jwk.New(t.PrivateKey)
	if err != nil {
		fmt.Printf("failed to create symmetric key: %s\n", err)
		return nil, err;
	}

	if _, ok := key.(jwk.RSAPrivateKey); !ok {
		fmt.Printf("expected jwk.SymmetricKey, got %T\n", key)
		return nil, err;
	}
	
	key.Set(jwk.KeyIDKey, t.JwkId)

	return key, nil
}
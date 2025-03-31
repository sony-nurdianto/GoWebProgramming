package encryption

import (
	"errors"
	"fmt"
	"os"
	t "time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

var (
	ErrSecretNotSet  error = errors.New("SECRET is not set")
	ErrDecryptFailed error = errors.New("invalid token failed to decrypt")
)

func secret() ([]byte, error) {
	secretKey := make([]byte, chacha20poly1305.KeySize)
	secret := os.Getenv("SECRET")
	if secret == "" {
		return nil, ErrSecretNotSet
	}

	copy(secretKey, secret)
	return secretKey, nil
}

func CreateWebToken(subject string) (string, error) {
	fmt.Println(subject)
	secretKey, err := secret()
	if err != nil {
		return "", err
	}

	jsonToken := paseto.JSONToken{
		Expiration: t.Now().Add(1 * t.Hour),
		Subject:    subject,
		Issuer:     "chitchat",
	}

	token, err := paseto.NewV2().Encrypt(secretKey, jsonToken, nil)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyWebToken(token string) (paseto.JSONToken, error) {
	var newJsonToken paseto.JSONToken

	secretKey, err := secret()
	if err != nil {
		return newJsonToken, err
	}

	if err := paseto.NewV2().Decrypt(token, secretKey, &newJsonToken, nil); err != nil {
		return newJsonToken, err
	}

	if newJsonToken.Expiration.Before(t.Now()) {
		return newJsonToken, ErrDecryptFailed
	}

	return newJsonToken, nil
}

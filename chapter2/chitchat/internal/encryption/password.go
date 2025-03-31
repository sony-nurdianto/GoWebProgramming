package encryption

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"

	"golang.org/x/crypto/argon2"
)

const (
	time    = 1
	memory  = 64 * 1024
	threads = 4
	keyLen  = 32
)

func GenerateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func HashPassword(password string) (string, error) {
	salt, err := GenerateSalt(16)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	saltStr := base64.StdEncoding.EncodeToString(salt)
	hashStr := base64.StdEncoding.EncodeToString(hash)

	return fmt.Sprintf("%s%s", hashStr, saltStr), nil
}

func VerifyPassword(password, hashPassword string) (bool, error) {
	salt, err := base64.StdEncoding.DecodeString(hashPassword[44:])
	if err != nil {
		log.Println("Error decoding salt:", err)
		return false, err
	}

	expectedHash, err := base64.StdEncoding.DecodeString(hashPassword[:44])
	if err != nil {
		log.Println("Error decoding hash:", err)
		return false, err
	}

	newHash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	return string(newHash) == string(expectedHash), nil
}

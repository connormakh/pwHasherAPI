package utils

import (
	"crypto/sha512"
	"encoding/base64"
)

func Hash(password string) string {
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hashed := hasher.Sum(nil)
	return base64.StdEncoding.EncodeToString(hashed)
}

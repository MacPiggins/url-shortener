package tokeniser

import (
	"crypto/sha256"
	"url-shortener/pkg/Base63"
)

func GenerateToken(link string) string {
	hash := sha256.Sum256([]byte(link))
	bytes := hash[:]
	token := Base63.ConvertToBase63(bytes)
	token = token[:10]
	return token
}
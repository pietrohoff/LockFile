package encryption

import (
	"crypto/sha256"
)

// Função utilitária para gerar uma chave segura a partir de uma senha
func generateKey(password string) []byte {
	hash := sha256.Sum256([]byte(password))
	return hash[:]
}

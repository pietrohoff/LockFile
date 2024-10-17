package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
	"path/filepath"
)

// Função de criptografia do arquivo
func EncryptFile(filePath string, password string) error {
	// Gera a chave de criptografia
	key := generateKey(password)

	// Abre o arquivo para leitura
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Cria o diretório "files/encrypted" caso ele não exista
	if _, err := os.Stat("files/encrypted"); os.IsNotExist(err) {
		os.MkdirAll("files/encrypted", os.ModePerm)
	}

	// Extrai o nome base do arquivo original e define o caminho de saída
	encryptedFileName := filepath.Base(filePath) + ".enc"
	outputFile, err := os.Create("files/encrypted/" + encryptedFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Inicializa o bloco de cifragem AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Gera um IV (Initialization Vector) aleatório
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	// Escreve o IV no início do arquivo criptografado
	if _, err := outputFile.Write(iv); err != nil {
		return err
	}

	// Criptografa o conteúdo do arquivo usando AES-256
	stream := cipher.NewCFBEncrypter(block, iv)
	writer := &cipher.StreamWriter{S: stream, W: outputFile}
	if _, err := io.Copy(writer, file); err != nil {
		return err
	}

	return nil
}

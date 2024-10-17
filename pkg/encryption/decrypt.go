package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"
	"path/filepath"
)

// Função de descriptografia do arquivo
func DecryptFile(filePath string, password string) error {
	// Gera a chave de descriptografia
	key := generateKey(password)

	// Abre o arquivo criptografado para leitura
	encryptedFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer encryptedFile.Close()

	// Cria o diretório "files/decrypted" caso ele não exista
	if _, err := os.Stat("files/decrypted"); os.IsNotExist(err) {
		os.MkdirAll("files/decrypted", os.ModePerm)
	}

	// Define o caminho de saída do arquivo descriptografado
	decryptedFileName := filepath.Base(filePath)
	decryptedFileName = decryptedFileName[:len(decryptedFileName)-4] // Remove ".enc" do nome
	outputFile, err := os.Create("files/decrypted/" + decryptedFileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Lê o IV (Initialization Vector) do início do arquivo
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(encryptedFile, iv); err != nil {
		return err
	}

	// Inicializa o bloco de decifragem AES
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	// Descriptografa o conteúdo do arquivo
	stream := cipher.NewCFBDecrypter(block, iv)
	reader := &cipher.StreamReader{S: stream, R: encryptedFile}
	if _, err := io.Copy(outputFile, reader); err != nil {
		return err
	}

	return nil
}

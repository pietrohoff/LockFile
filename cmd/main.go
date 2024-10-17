package main

import (
	"LockFile/pkg/encryption"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: go run main.go <caminho_do_arquivo> <senha> [--decrypt]")
		return
	}

	filePath := os.Args[1]
	password := os.Args[2]
	decrypt := len(os.Args) > 3 && os.Args[3] == "--decrypt"

	var err error
	if decrypt {
		err = encryption.DecryptFile(filePath, password)
		if err != nil {
			fmt.Println("Erro ao descriptografar o arquivo:", err)
			return
		}
		fmt.Println("Arquivo descriptografado com sucesso!")
	} else {
		err = encryption.EncryptFile(filePath, password)
		if err != nil {
			fmt.Println("Erro ao criptografar o arquivo:", err)
			return
		}
		fmt.Println("Arquivo criptografado com sucesso!")
	}
}

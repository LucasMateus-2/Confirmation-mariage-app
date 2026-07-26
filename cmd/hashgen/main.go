package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("uso: go run cmd/hashgen/main.go <sua_senha>")
		os.Exit(1)
	}

	password := os.Args[1]

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("erro ao gerar hash:", err)
		os.Exit(1)
	}

	fmt.Println("Senha:", password)
	fmt.Println("Hash: ", string(hash))
}

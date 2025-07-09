package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	var length int
	fmt.Print("Введите длину пароля: ")
	fmt.Scan(&length)

	password := generatePassword(length)
	fmt.Println("Сгенерированный пароль:", password)
}

func generatePassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789" +
		"!@#$%^&*()-_=+[]{}<>?/|"

	result := make([]byte, length)

	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		result[i] = charset[num.Int64()]
	}

	return string(result)
}

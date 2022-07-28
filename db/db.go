package db

import "fmt"

type Database struct {
	Password   string
	Plaintext  string
	Ciphertext string
}

var Db = &Database{}

func ReadPlaintext() {
	fmt.Println("Input Text to Encrypt:")
	fmt.Scanln(&Db.Plaintext)
	fmt.Println("")
	fmt.Println("Plaintext input in bytes:")
	fmt.Println([]byte(Db.Plaintext))
	fmt.Println("")
}

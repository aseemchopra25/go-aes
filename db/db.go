package db

import "fmt"

type Database struct {
	Password   string
	Plaintext  string
	Ciphertext string
}

var Db = &Database{}

func ReadPlaintext() {
	fmt.Scanln(&Db.Plaintext)
}

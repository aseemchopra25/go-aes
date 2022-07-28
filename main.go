package main

import (
	"github.com/aseemchopra25/go-aes/db"
	"github.com/aseemchopra25/go-aes/decrypt"
	"github.com/aseemchopra25/go-aes/encrypt"
	"github.com/aseemchopra25/go-aes/keygen"
)

func main() {
	db.ReadPlaintext()                // 1. Read Plaintext Input
	keygen.KeyGen()                   // 2. Generate random 16 byte key
	encrypt.Encrypt(db.Db.Plaintext)  // 3. Encrypt plaintext
	decrypt.Decrypt(db.Db.Ciphertext) // 4. Decrypt Ciphertext
}

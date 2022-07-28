package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"

	"github.com/aseemchopra25/go-aes/db"
)

func handle(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Decrypt(encrypted string) {
	key, err := hex.DecodeString(db.Db.Password)
	handle(err)
	enc, err := hex.DecodeString(encrypted)
	handle(err)
	block, err := aes.NewCipher(key)
	handle(err)
	aesGCM, err := cipher.NewGCM(block)
	handle(err)
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	dec, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	handle(err)
	fmt.Println("Decrypted Output in bytes:")
	fmt.Println(dec)
	fmt.Println("")
	fmt.Println("Decrypted Output in UTF-8:")
	fmt.Println(string(dec))
	fmt.Println("")
}

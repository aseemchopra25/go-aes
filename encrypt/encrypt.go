package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/aseemchopra25/go-aes/db"
)

func handle(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func Encrypt() {
	key, err := hex.DecodeString(db.Db.Password) // retrieve bytes
	handle(err)
	block, err := aes.NewCipher(key)
	handle(err)
	aesGCM, err := cipher.NewGCM(block)
	handle(err)
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := rand.Reader.Read(nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := aesGCM.Seal(nonce, nonce, []byte(db.Db.Plaintext), nil)
	fmt.Println("Ciphertext Generated:")
	fmt.Println(ciphertext)
	fmt.Println("")
	db.Db.Ciphertext = hex.EncodeToString(ciphertext)

}

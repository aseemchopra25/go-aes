package keygen

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/aseemchopra25/go-aes/db"
)

func KeyGen() {
	b := make([]byte, 16)                          // 16 byte key for AES-128
	if _, err := rand.Reader.Read(b); err != nil { //switched math/rand to crypto/rand
		panic(err.Error())
	}
	fmt.Println("128 Bit Key:")
	fmt.Println(b)
	fmt.Println("")
	db.Db.Password = hex.EncodeToString(b)
}

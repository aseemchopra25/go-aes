package keygen

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/aseemchopra25/go-aes/db"
)

func KeyGen() {
	b := make([]byte, 16)                          // 16 bit key for AES-128
	if _, err := rand.Reader.Read(b); err != nil { //switched math/rand to crypto/rand
		panic(err.Error())
	}
	db.Db.Password = hex.EncodeToString(b)
}

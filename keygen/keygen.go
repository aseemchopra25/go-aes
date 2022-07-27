package keygen

import (
	"encoding/hex"
	"fmt"
	"math/rand"
)

func KeyGen() {
	b := make([]byte, 16) // 16 bit key for AES-128
	if _, err := rand.Read(b); err != nil {
		panic(err.Error())
	}
	key := hex.EncodeToString(b)
	fmt.Println(key)
}

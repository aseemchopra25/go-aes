package keygen

import (
	"fmt"
	"math/rand"
)

func KeyGen() {
	b := make([]byte, 16) // 16 bit key for AES-128
	if _, err := rand.Read(b); err != nil {
		panic(err.Error())
	}
	fmt.Println(b)
}

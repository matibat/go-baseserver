package usables

import (
	"math/big"
	"math/rand"
	"time"
)

const base int64 = 64

var basebig = big.NewInt(64)

var digits = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ._")

// GetRandBytes -> Genera un string
func GetRandBytes(length int) []byte {
	var output []byte
	var char byte
	var src = rand.NewSource(int64(time.Now().UnixNano()))
	src.Seed(int64(time.Now().UnixNano()))
	var rnd = rand.New(src)
	for i := 0; i < length; i++ {
		char = byte(rnd.Int())
		output = append(output, char)
	}
	return output
}


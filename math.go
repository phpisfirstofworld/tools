package tools

import (
	"crypto/rand"
	"math/big"
)

// Mt_rand 老版本兼容函数
func Mt_rand(min, max int64) int64 {

	n, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))

	return n.Int64() + min
}

func MtRand(min, max int64) int64 {

	n, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))

	return n.Int64() + min
}

package tools

import (
	"math/rand"
	"time"
)

//老版本兼容函数
func Mt_rand(min, max int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min+1) + min
}

//生成固定区间的随机数
func MtRand(min, max int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min+1) + min
}

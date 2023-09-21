package tools

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HashHmac(data []byte, key []byte, binary bool) string {

	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(data))
	h := mac.Sum(nil)
	if binary {
		return string(h)
	}
	return hex.EncodeToString(h)

}

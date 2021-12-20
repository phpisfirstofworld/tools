package secret

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

func AesEncryptCFB(origData []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {

		return []byte{}, err
	}
	encrypted := make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {

		return []byte{}, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)

	return []byte(hex.EncodeToString(encrypted)), nil
}
func AesDecryptCFB(encrypted []byte, key []byte) ([]byte, error) {

	encrypted, err := hex.DecodeString(string(encrypted))

	if err != nil {

		return []byte{}, err
	}

	block, err := aes.NewCipher(key)

	if err != nil {

		return []byte{}, err
	}

	if len(encrypted) < aes.BlockSize {
		//panic("ciphertext too short")
		return []byte{}, errors.New("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)

	return encrypted, nil

}

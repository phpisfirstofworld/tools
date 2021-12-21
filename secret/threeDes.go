package secret

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/PeterYangs/tools"
)

type Types int

const (
	Hex      Types = 1
	Base64   Types = 2
	Original Types = 3
)

type desSecret struct {
}

type desPassword struct {
	original []byte
}

func NewDesPassword(original []byte) *desPassword {

	return &desPassword{original: original}
}

func (dp *desPassword) ToBase64() []byte {

	return []byte(base64.StdEncoding.EncodeToString(dp.original))
}

func (dp *desPassword) ToHex() []byte {

	return []byte(hex.EncodeToString(dp.original))
}

func (dp *desPassword) ToOriginal() []byte {

	return dp.original
}

func NewDes() *desSecret {

	return &desSecret{}
}

// GenerateKey 生成key
func (d *desSecret) GenerateKey() []byte {

	str := "1234567890qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"

	key := ""
	for i := 0; i < 24; i++ {

		index := tools.MtRand(0, int64(len(str)-1))

		key += str[index : index+1]

	}

	return []byte(key)

}

//填充字符串（末尾）
func (d *desSecret) paddingText(str []byte, blockSize int) []byte {
	//需要填充的数据长度
	paddingCount := blockSize - len(str)%blockSize
	//填充数据为：paddingCount ,填充的值为：paddingCount
	paddingStr := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	newPaddingStr := append(str, paddingStr...)
	//fmt.Println(newPaddingStr)
	return newPaddingStr
}

//去掉字符（末尾）
func (d *desSecret) unPaddingText(str []byte) ([]byte, error) {
	n := len(str)
	count := int(str[n-1])
	//fmt.Println("xxxxxxxxxxxx:", n-count)

	if n-count < 0 || n-count > len(str) {

		return []byte{}, errors.New("截取长度异常")
	}

	newPaddingText := str[:n-count]
	return newPaddingText, nil
}

func (d *desSecret) Encyptog3DES(src, key []byte) (*desPassword, error) {
	//des包下的三次加密接口
	block, err := des.NewTripleDESCipher(key)

	if err != nil {

		return nil, err
	}

	src = d.paddingText(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(src, src)

	return NewDesPassword(src), nil
}
func (d *desSecret) Decrptog3DES(src, key []byte, types Types) ([]byte, error) {

	var err error

	switch types {

	case Hex:

		src, err = hex.DecodeString(string(src))

	case Base64:
		src, err = base64.StdEncoding.DecodeString(string(src))

	case Original:

	default:

	}

	if err != nil {

		return []byte{}, err
	}

	block, err := des.NewTripleDESCipher(key)
	if err != nil {

		return []byte{}, err
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(src, src)
	src, err = d.unPaddingText(src)

	if err != nil {

		return []byte{}, err
	}

	return src, nil
}

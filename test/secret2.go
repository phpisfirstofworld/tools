package main

import (
	"fmt"
	"github.com/PeterYangs/tools/secret"
)

func main() {

	d := secret.NewDes()

	source := "hello world"
	fmt.Println("原字符：", source)

	key := d.GenerateKey() //24位

	//加密
	code, err := d.Encyptog3DES([]byte(source), []byte(key))

	if err != nil {

		fmt.Println("加密错误", err)

		return
	}

	fmt.Println("密文：", code.ToBase64())

	//解密
	real, err := d.Decrptog3DES([]byte(code.ToBase64()), []byte(key), secret.Base64)
	//
	if err != nil {

		fmt.Println("解密错误", err)

		return
	}

	fmt.Println("解密：", string(real))

}

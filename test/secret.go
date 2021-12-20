package main

import (
	"fmt"
	"github.com/PeterYangs/tools/secret"
)

func main() {

	source := "hello world"
	fmt.Println("原字符：", source)
	key := "4179cdded7fbc8f3936a4494cb7dc46b" //16位

	code, err := secret.AesEncryptCFB([]byte(source), []byte(key))

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println("密文：", string(code))

	real, err := secret.AesDecryptCFB(code, []byte("4179cdded7fbc8f3936a4494cb7dc46c"))

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println("解密：", string(real))

}

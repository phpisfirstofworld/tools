package main

import (
	"fmt"
	"github.com/PeterYangs/tools/file"
)

func main() {

	f, e := file.OpenFileWithAPPEND("xx.txt")

	if e != nil {

		fmt.Println(e)

		return
	}

	file.ReadFileToByte("array.go", 1024, func(bytes []byte) {

		//fmt.Println(bytes)

		f.Write(bytes)

	})

}

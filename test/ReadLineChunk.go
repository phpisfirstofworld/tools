package main

import (
	"fmt"
	"github.com/PeterYangs/tools/file"
)

func main() {

	file.ReadLineChunk("123.txt", 0, 3, func(list []string) {

		fmt.Println(list)

	})

}

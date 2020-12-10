package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	re, _ := tools.GetWithString("https://github.com")

	fmt.Println(re)
}

package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	re, _ := tools.GetToString("https://github.com", tools.HttpSetting{})

	fmt.Println(re)
	fmt.Println(re)
}

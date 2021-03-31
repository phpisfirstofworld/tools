package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	str := "我尼玛"

	s := tools.SubStr(str, 1, 1)
	//s:=tools.SubStr(str,1,-1)

	fmt.Println(s)
	fmt.Println(str)
}

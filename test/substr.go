package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	str := "我尼玛asdasd我萨达奥斯卡就&"

	//s_:=[]rune(str)
	//
	//fmt.Println(s_[1:2])

	s := tools.SubStr(str, 0, -2)
	//s:=tools.SubStr(str,1,-1)

	fmt.Println(s)
	fmt.Println(str)
}

package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	name := "1212.png"

	f1, _ := tools.GetExtensionName(name)

	fmt.Println(f1)

	name2 := "sdasd.sadsad.asdsad.asd.jpg"

	f2, _ := tools.GetExtensionName(name2)

	fmt.Println(f2)

	name3 := "文.文as打卡机...png"

	f3, _ := tools.GetExtensionName(name3)

	fmt.Println(f3)

	name4 := "https://item-shopping.c.yimg.jp/i/n/sakuranokoi_udj008_1"

	f4, _ := tools.GetExtensionName(name4)

	fmt.Println(f4)
}

package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	name := "1212.png"

	f1 := tools.GetExtensionName(name)

	fmt.Println(f1)

	name2 := "sdasd.sadsad.asdsad.asd.jpg"

	f2 := tools.GetExtensionName(name2)

	fmt.Println(f2)

	name3 := "文.文as打卡机...png"

	f3 := tools.GetExtensionName(name3)

	fmt.Println(f3)

}

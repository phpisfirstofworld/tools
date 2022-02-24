package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	str := "asdkhkyyjasdashkjd"

	l := tools.StrPos(str, "yy")

	fmt.Println(l)

}

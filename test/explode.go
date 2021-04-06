package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	str := "1,2,3"

	fmt.Println(tools.Explode(",", str))

}

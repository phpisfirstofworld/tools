package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	arr := []string{"1", "2", "3"}

	fmt.Println(tools.Implode("-", arr))

}

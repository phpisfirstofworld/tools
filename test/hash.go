package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	h := tools.HashHmac([]byte("132"), []byte("456"), false)

	fmt.Println(h)

}

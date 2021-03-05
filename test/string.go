package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	html, _ := tools.GetToString("https://www.duote.com/", tools.HttpSetting{})

	fmt.Println(tools.IsGBK([]byte(html)))

}

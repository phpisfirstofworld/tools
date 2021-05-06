package main

import (
	"fmt"
	"github.com/PeterYangs/tools/http"
)

func main() {

	client := http.Client(http.HttpSetting{})

	html, err := client.GetToString("https://www.baidu.com")

	if err != nil {

		fmt.Println(err)

		return

	}
	fmt.Println(html)

}

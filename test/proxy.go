package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	url := "https://www.google.com/"
	//url := "https://www.baidu.com"

	html, err := tools.GetToString(url, tools.HttpSetting{
		//ProxyAddress: "http://127.0.0.1:4780",
		ProxyAddress: "socks5://127.0.0.1:4781",
	})

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(html)

}

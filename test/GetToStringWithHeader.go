package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	url := "https://www.4399.com"

	_, header, err := tools.GetToStringWithHeader(url, tools.HttpSetting{})

	if err != nil {

		fmt.Println(err)

		return
	}

	//fmt.Println(html)
	fmt.Println(header["Set-Cookie"])

}

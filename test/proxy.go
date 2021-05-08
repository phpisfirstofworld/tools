package main

import (
	"fmt"
	"github.com/PeterYangs/tools/http"
)

func main() {

	client := http.Client()

	client.SetProxyAddress("http://127.0.0.1:4780")

	html, err := client.Request().GetToString("https://www.google.com/")

	if err != nil {

		fmt.Println(err)
	}

	fmt.Println(html)

}

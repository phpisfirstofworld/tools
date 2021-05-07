package main

import (
	"fmt"
	"github.com/PeterYangs/tools/http"
	"time"
)

func main() {

	client := http.Client(http.Setting{})

	header := map[string]string{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"}

	str, err := client.SetTimeout(1 * time.Second).Request().SetHeader(header).GetToString("https://www.google.com/")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Print(str)

}

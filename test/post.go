package main

import (
	"fmt"
	"github.com/PeterYangs/tools/http"
	"time"
)

func main() {

	client := http.Client(http.Setting{})

	p := map[string]interface{}{"name": []string{"123", "456"}, "age": 1, "nickname": "123"}

	str, err := client.SetTimeout(1 * time.Second).Request().SetParameter(p).GetToString("https://www.google.com/")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Print(str)

}

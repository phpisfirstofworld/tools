package main

import (
	"fmt"
	"github.com/PeterYangs/tools/http"
)

func main() {

	header := map[string]string{"user-agent": "Iphone100"}

	p := map[string]interface{}{"name": []string{"123", "456"}, "age": 1, "nickname": "123"}

	client := http.Client()

	re, err := client.Request().SetHeader(header).SetParameter(p).PostToString("http://list.com/pass/post.php")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(re)

}

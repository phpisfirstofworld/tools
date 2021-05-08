package main

import (
	"fmt"
	"github.com/PeterYangs/tools/http"
)

func main() {

	header := map[string]string{"user-agent": "Iphone100"}

	p := map[string]interface{}{"name": []string{"123", "456"}, "age": 1, "nickname": "123"}
	//p2 := map[string]interface{}{"nickname": "456"}

	client := http.Client()

	client.SetHeader(header)

	re, err := client.Request().SetParameter(p).SetHeader(map[string]string{"user-agent": "Iphone10"}).PostToString("http://list.com/pass/header.php")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(re)

	//re, err =req.PostToString("http://list.com/pass/post.php")
	//
	//if err != nil {
	//
	//	fmt.Println(err)
	//
	//	return
	//}
	//
	//fmt.Println(re)

}

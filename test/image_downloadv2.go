package main

import (
	"fmt"
	"github.com/PeterYangs/tools/http"
)

func main() {

	client := http.Client()

	//p := map[string]interface{}{"name": []string{"123", "456"}, "age": 1, "nickname": "123"}

	err := client.Request().DownloadImage("https://down.925g.com/upload/cms/20201225/1745/e1162c3215414a581e43110c5f8c886c.png", "yy.png")
	//err := client.Request().DownloadFile("https://down.925g.com/upload/cms/20201225/1745/e1162c3215414a581e43110c5f8c886c.png","yy.png")

	if err != nil {

		fmt.Println(err)

		return
	}

	//fmt.Print(str)

}

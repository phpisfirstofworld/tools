package main

import (
	"fmt"
	"github.com/PeterYangs/tools/http"
)

func main() {

	//获取客户端
	client := http.Client()

	//携带参数
	p := map[string]interface{}{
		"name":     []string{"123", "456"},
		"age":      1,
		"nickname": "123",
		"form": map[string]interface{}{
			"one":   "1",
			"two":   "2",
			"three": []string{"123", "456"},
			"four": map[string]interface{}{
				"one": "1",
				"two": "2",
			},
		},
	}

	for i := 0; i < 2; i++ {

		str, err := client.Request().SetParameter(p).GetToString("http://list.com/pass/get.php")

		if err != nil {

			fmt.Println(err)

			return
		}

		fmt.Println(str)

	}

}

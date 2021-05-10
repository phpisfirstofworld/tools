package main

import (
	"fmt"
	"github.com/PeterYangs/tools/http"
)

func main() {

	client := http.Client()

	i, err := client.Request().SetReTryTimes(3).GetToString("https://xxxccaacasdad.com")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(i)

}

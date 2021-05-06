package main

import (
	"fmt"
	"github.com/PeterYangs/tools/http"
)

var client *http.HttpClient

func main() {

	client = http.Client(http.HttpSetting{})

	max := make(chan int, 30)

	for {

		max <- 1

		go run(max)

	}

}

func run(max chan int) {

	defer func(maxs chan int) {

		<-maxs

	}(max)

	h, err := client.GetToString("https://www.youxi369.com/")

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println(h)

}

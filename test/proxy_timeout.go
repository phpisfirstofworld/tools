package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
	uuid "github.com/satori/go.uuid"
)

func main() {

	var c = make(chan int, 50)

	for i := 0; i < 3000; i++ {

		c <- 1

		go func(c chan int) {

			defer func() {

				<-c
			}()

			err := tools.DownloadImage(
				"https://down.925g.com/upload/cms/20210326/1445/e1e9c92c6a7be74b62df7c6a0a658302.jpg",
				"img/"+uuid.NewV4().String()+".png",
				tools.HttpSetting{},
			)

			if err != nil {

				fmt.Println(err)
			}

		}(c)

	}

}

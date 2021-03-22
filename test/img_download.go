package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	img := "https://item-shopping.c.yimg.jp/i/n/sakuranokoi_udj008_1"

	//err := tools.DownloadFile(img, "image.png", tools.HttpSetting{})
	err := tools.DownloadImage(img, "image.png", tools.HttpSetting{})

	fmt.Println(err)

	//tools.DownloadFile(img, "image.png")

}

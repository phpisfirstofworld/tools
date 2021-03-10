package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
)

func main() {

	img := "https://down.925g.com/upload/cms/20210301/1655/193cf532be3661a57d24dc8e100ffecc.png"

	err := tools.DownloadFile(img, "image.png")
	//err:=tools.DownloadImage(img, "image.png")

	fmt.Println(err)

	//tools.DownloadFile(img, "image.png")

}

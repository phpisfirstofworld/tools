package main

import "github.com/PeterYangs/tools"

func main() {

	img := "https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1607593445045&di=ac39875471c30a8c66597cba29db37ad&imgtype=0&src=http%3A%2F%2Fattachments.gfan.com%2Fforum%2F201503%2F19%2F211608ztcq7higicydxhsy.jpg"

	tools.DownloadImage(img, "image.png")

}

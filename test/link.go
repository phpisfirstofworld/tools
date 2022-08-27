package main

import (
	"fmt"
	"github.com/PeterYangs/tools/link"
)

func main() {

	//fmt.Println(filepath.Dir("www.xyzs.com/public/css/style.css"))

	//fmt.Println(link.GetCompleteLink("https://www.xyzs.com/public/css/style.css", "../img/bgrat.png"))

	fmt.Println(link.GetCompleteLink("https://www.522gg.com/pc/css/iconfont.css", "iconfont.woff2?t=1625723377069"))

	//u, _ := url.Parse("https://www.522gg.com/pc/css/iconfont.css")
	//
	//fmt.Println()

}

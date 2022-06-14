package main

import (
	"fmt"
	"github.com/PeterYangs/tools/link"
)

func main() {

	//fmt.Println(filepath.Dir("www.xyzs.com/public/css/style.css"))

	fmt.Println(link.GetCompleteLink("https://www.xyzs.com/public/css/style.css", "../img/bgrat.png"))

}

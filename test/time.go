package main

import (
	"fmt"
	"github.com/PeterYangs/tools"
	"time"
)

func main() {

	fmt.Println(tools.Date("Y-m-d", time.Now().Unix()))
	fmt.Println(tools.Date("Y-m-d H:i:s", time.Now().Unix()))
	fmt.Println(tools.Date("Y-m-d H:i", time.Now().Unix()))
	fmt.Println(tools.Date("Y", time.Now().Unix()))
	fmt.Println(tools.Date("Ym", time.Now().Unix()))
	fmt.Println(tools.Date("Ymd", time.Now().Unix()))
	fmt.Println(tools.Date("H", time.Now().Unix()))
	fmt.Println(tools.Date("Hi", time.Now().Unix()))

	fmt.Println(tools.StrToTime("2020/12/12"))
	fmt.Println(tools.StrToTime("2020-12-12"))
	fmt.Println(tools.StrToTime("2020-12-12 11:32:00"))
	fmt.Println(tools.StrToTime("2020/12/12 11:32:00"))

}

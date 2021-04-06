package main

import "github.com/PeterYangs/tools"

func main() {

	tools.GetToString("https://www.baidu.com", tools.HttpSetting{TimeOut: 10})

}

package main

import "github.com/PeterYangs/tools"

func main() {

	header := map[string]string{"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36"}

	tools.GetToString("https://www.baidu.com", tools.HttpSetting{Header: header})

}

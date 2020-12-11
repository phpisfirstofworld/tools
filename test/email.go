package main

import (
	"github.com/PeterYangs/tools"
)

func main() {

	tools.SendEmail(
		"904801074@qq.com",
		[]string{"904801074@qq.com"},
		"title",
		"<h1>hello</h1>",
		"smtp.qq.com",
		465,
		"******",
	)

}

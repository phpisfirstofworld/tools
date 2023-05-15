package main

import (
	"fmt"
	"github.com/PeterYangs/tools/file"
)

func main() {

	//获取目标文件夹下的所有文件（包含子目录，返回false则不继续遍历）
	file.GetDirList("./", func(path string) bool {

		fmt.Println(path)

		return true
	})

}

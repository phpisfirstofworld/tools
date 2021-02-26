package tools

import (
	"io/ioutil"
	"os"
	"strings"
)

func WriteLine(path string, data string) {

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {

		//fmt.Println(err)

		panic(err)
	}

	defer f.Close()

	_, err = f.Write([]byte(data + "\n"))

	if err != nil {

		//fmt.Println(err)

		panic(err)
	}

}

func DeleteFile(path string) error {

	err := os.Remove(path)

	if err != nil {

		//fmt.Println("文件删除错误,", err)

		return err
	}

	return nil

}

//读取文件
func ReadFile(path string) (string, error) {

	data, err := ioutil.ReadFile(path)

	return string(data), err
}

//获取文件名拓展名
func GetExtensionName(fileName string) string {

	index := strings.LastIndexFunc(fileName, func(r rune) bool {

		if string(r) == "." {

			return true
		}

		return false
	})

	return fileName[index+1:]
}

package tools

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

func WriteLine(path string, data string) {

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

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

//删除文件
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
func GetExtensionName(fileName string) (string, error) {

	index := strings.LastIndexFunc(fileName, func(r rune) bool {

		if string(r) == "." {

			return true
		}

		return false
	})

	ex := fileName[index+1:]

	f := strings.Contains(ex, "/")

	if f {

		return "", errors.New("extension is error")
	}

	return ex, nil
}

//创建多级文件夹
func MkDirDepth(path string) error {

	array := Explode("/", path)

	currentPath := ""

	for _, v := range array {

		if v == "" {

			continue
		}

		currentPath = currentPath + v + "/"

		err := os.Mkdir(currentPath, 0644)

		if err != nil && err.Error() != "mkdir "+currentPath+": Cannot create a file when that file already exists." {

			return err
		}

	}

	return nil
}

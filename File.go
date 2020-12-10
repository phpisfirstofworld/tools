package tools

import (
	"fmt"
	"io/ioutil"
	"os"
)

func WriteLine(path string, data string) {

	//panic(path+"-------------"+data)

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	if err != nil {

		//fmt.Println(err)

		panic(err)
	}

	//f.Write([]byte(data + "\n"))

	defer f.Close()

	_, err = f.Write([]byte(data + "\n"))

	if err != nil {

		//fmt.Println(err)

		panic(err)
	}

}

func DeleteFile(path string) {

	err := os.Remove(path)

	if err != nil {

		fmt.Println("文件删除错误,", err)

	}

}

//读取文件
func ReadFile(path string) (string, error) {

	data, err := ioutil.ReadFile(path)

	return string(data), err
}

package file

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// ReadFile 读取整个文件
func ReadFile(path string) (string, error) {

	data, err := ioutil.ReadFile(path)

	return string(data), err
}

// ReadFileToByte 分块读取文件，第三个参数是每块的回调
func ReadFileToByte(path string, bufSize int, callback func([]byte)) error {

	f, err := os.Open(path)

	if err != nil {

		fmt.Println(err)

		return err
	}

	defer f.Close()

	for {

		buf := make([]byte, bufSize)

		l, e := f.Read(buf)

		if e != nil {

			//读取完毕
			if err == io.EOF {

				return nil
			}

			return e

		}

		callback(buf[:l])

	}

}

// OpenFileWithAPPEND 追加打开文件（别忘了关闭文件）
func OpenFileWithAPPEND(path string) (*os.File, error) {

	f, e := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_EXCL, 0644)

	if e != nil {

		return f, e
	}

	return f, nil
}

// GetExtensionName 获取文件名拓展名
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

// DeleteFile 删除文件
func DeleteFile(path string) error {

	err := os.Remove(path)

	if err != nil {

		//fmt.Println("文件删除错误,", err)

		return err
	}

	return nil

}

// WriteLine 逐行写入
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

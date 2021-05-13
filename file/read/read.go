package read

import (
	"bytes"
	"io"
	"os"
)

type File struct {
	file      *os.File
	openError error
}

// Open 打开文件
func Open(path string) *File {

	f, err := os.Open(path)

	if err != nil {

		return &File{file: f, openError: err}

	}

	return &File{file: f, openError: nil}

}

//读取文件
func (file *File) Read() ([]byte, error) {

	if file.openError != nil {

		return []byte{}, file.openError
	}

	defer file.file.Close()

	//存放结果
	result := bytes.NewBuffer(nil)

	defer result.Reset()

	//缓冲字节
	buf := make([]byte, 1024)

	for {

		//读到缓冲，每次读取都会清空上次数据
		l, err := file.file.Read(buf)

		if err != nil {

			if err == io.EOF {

				break
			}

			return nil, err
		}
		//写入到结果
		result.Write(buf[:l])

	}
	return result.Bytes(), nil
}

// ReadBlock 分块读取文件
func (file *File) ReadBlock(bufSize int, callback func([]byte)) error {

	if file.openError != nil {

		return file.openError
	}

	defer file.file.Close()

	//缓冲字节
	buf := make([]byte, bufSize)

	for {

		//读到缓冲，每次读取都会清空上次数据
		l, err := file.file.Read(buf)

		if err != nil {

			if err == io.EOF {

				break
			}

		}

		callback(buf[:l])

	}

	return nil
}

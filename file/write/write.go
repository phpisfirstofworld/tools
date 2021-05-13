package write

import (
	"os"
)

type File struct {
	file      *os.File
	openError error
}

// Open 打开文件，文件不存在创建，文件已存在清空
func Open(path string) *File {

	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {

		return &File{file: f, openError: err}

	}

	return &File{file: f, openError: nil}

}

// OpenAppend 追加模式打开文件，文件不存在创建，文件已存在清空
func OpenAppend(path string) *File {

	f, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0644)

	if err != nil {

		return &File{file: f, openError: err}

	}

	return &File{file: f, openError: nil}

}

// OpenMustNotExist 打开文件，文件不存在创建，文件已存在就报错
func OpenMustNotExist(path string) *File {

	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL, 0644)

	if err != nil {

		return &File{file: f, openError: err}

	}

	return &File{file: f, openError: nil}
}

// OpenAppendMustNotExist 追加打开文件，文件不存在创建，文件已存在就报错
func OpenAppendMustNotExist(path string) *File {

	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_APPEND, 0644)

	if err != nil {

		return &File{file: f, openError: err}

	}

	return &File{file: f, openError: nil}
}

// GetFile 获取文件
func (file *File) GetFile() (*os.File, error) {

	if file.openError != nil {

		return nil, file.openError
	}

	return file.file, nil
}

func (file *File) Write(data []byte) error {

	if file.openError != nil {

		return file.openError
	}

	//defer file.file.Close()

	_, err := file.file.Write(data)

	return err
}

func (file *File) Close() error {

	return file.file.Close()

}

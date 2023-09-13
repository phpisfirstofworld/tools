package file

import (
	"bufio"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"os"
)

// Read 一次性读取
func Read(path string) ([]byte, error) {

	f, err := os.Open(path)

	if err != nil {

		return []byte{}, err
	}

	defer func() {

		_ = f.Close()

	}()

	return ioutil.ReadAll(f)
}

// ReadLine 逐行读取
func ReadLine(path string, callback func(line []byte)) error {

	f, err := os.Open(path)

	if err != nil {

		return err
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {

		l, _, e := r.ReadLine()

		if e != nil {

			if e != io.EOF {

				return e
			}

			break

		}

		callback(l)

	}

	return nil
}

// ReadLineWithBreak  逐行读取可退出
func ReadLineWithBreak(path string, callback func(line []byte) bool) error {

	f, err := os.Open(path)

	if err != nil {

		return err
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {

		l, _, e := r.ReadLine()

		if e != nil {

			if e != io.EOF {

				return e
			}

			break

		}

		ok := callback(l)

		if !ok {

			break
		}

	}

	return nil
}

func ReadLineWithCxt(cxt context.Context, path string, callback func(line []byte)) error {

	f, err := os.Open(path)

	if err != nil {

		return err
	}

	defer f.Close()

	r := bufio.NewReader(f)

	for {

		select {
		case <-cxt.Done():

			return errors.New("cxt cancel")

		default:

		}

		l, _, e := r.ReadLine()

		if e != nil {

			if e != io.EOF {

				return e
			}

			break

		}

		callback(l)

	}

	return nil
}

func ReadLineChunk(path string, offset int, size int, callback func(list []string)) error {

	f, err := os.Open(path)

	if err != nil {

		return err
	}

	defer f.Close()

	var c []string

	r := bufio.NewReader(f)

	currentOffset := 0

	for {

		l, _, e := r.ReadLine()

		if e != nil {

			if e != io.EOF {

				return e
			}

			if len(c) != 0 {

				callback(c)
			}

			break

		}

		if offset > currentOffset {

			currentOffset++

			continue
		}

		c = append(c, string(l))

		if len(c) >= size {

			callback(c)

			c = []string{}
		}

		currentOffset++
	}

	return nil
}

// Write 覆盖写
func Write(path string, content []byte) error {

	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)

	if err != nil {

		return err
	}

	defer f.Close()

	f.Write(content)

	return nil
}

// GetDirList 获取目标文件夹下的所有文件(包含子目录)
func GetDirList(dir string, fun func(path string) bool) error {

	return getDirList("", fun, dir)

}

func getDirList(dir string, fun func(path string) bool, sourceDir string) error {

	files, err := ioutil.ReadDir(sourceDir + If(dir == "", "", "/"+dir))

	if err != nil {

		return err

	}

	for _, file := range files {

		if !file.IsDir() {

			isGo := fun(If(dir == "", "", dir+"/") + file.Name())

			if isGo == false {

				break
			}

		} else {

			getDirList(If(dir == "", "", dir+"/")+file.Name(), fun, sourceDir)

		}

	}

	return nil

}

func If(condition bool, trueVal, falseVal string) string {

	if condition {
		return trueVal
	}
	return falseVal
}

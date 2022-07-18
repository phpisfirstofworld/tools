package file

import (
	"bufio"
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

func Write(path string, content []byte) error {

	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)

	if err != nil {

		return err
	}

	defer f.Close()

	f.Write(content)

	return nil
}

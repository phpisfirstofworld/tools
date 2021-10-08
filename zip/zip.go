package zip

import (
	zip2 "archive/zip"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type zip struct {
}

func NewZip() *zip {

	return &zip{}
}

type unzip struct {
	path string
}

func (z *zip) Open(path string) *unzip {

	return &unzip{
		path: path,
	}

}

func (u *unzip) Extract(toDir string) error {

	rr, err := zip2.OpenReader(u.path)

	defer rr.Close()

	if err != nil {

		return err
	}

	for _, k := range rr.Reader.File {

		decodeName := ""

		if k.Flags == 0 {
			//如果标致位是0  则是默认的本地编码   默认为gbk
			i := bytes.NewReader([]byte(k.Name))

			decoder := transform.NewReader(i, simplifiedchinese.GB18030.NewDecoder())
			content, _ := ioutil.ReadAll(decoder)
			decodeName = string(content)
		} else {
			//如果标志为是 1 << 11也就是 2048  则是utf-8编码
			decodeName = k.Name
		}

		//文件夹创建
		if k.FileInfo().IsDir() {
			err := os.MkdirAll(toDir+"/"+decodeName, 0755)
			if err != nil {

				return err

			}
			continue
		}

		r, err := k.Open()
		if err != nil {

			return err
		}

		NewFile, err := os.Create(toDir + "/" + decodeName)
		if err != nil {

			return err
		}
		io.Copy(NewFile, r)
		NewFile.Close()

	}

	return nil
}

type pzip struct {
	path string
}

func (z *zip) Create(path string) *pzip {

	return &pzip{
		path: path,
	}
}

func (p *pzip) PackDir(dir string) error {

	// 创建：zip文件
	zipfile, err := os.OpenFile(p.path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0644)

	if err != nil {

		return err
	}

	defer zipfile.Close()

	// 打开：zip文件
	archive := zip2.NewWriter(zipfile)

	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(dir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == dir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip2.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, dir+`\`)

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip2.Deflate
		}

		// 创建：压缩包头部信息
		writer, _ := archive.CreateHeader(header)
		if !info.IsDir() {
			file, _ := os.Open(path)
			defer file.Close()
			io.Copy(writer, file)
		}
		return nil
	})

	return nil
}

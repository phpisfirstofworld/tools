package tools

import (
	"compress/gzip"
	_ "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

//超时时间
const HttpTimeOut = 10 * time.Second

//请求底层函数
func Query(url string, method string, parameter string, header map[string]string) (*http.Response, error) {

	client := http.Client{}

	client.Timeout = HttpTimeOut

	req, err := http.NewRequest(method, url, strings.NewReader(parameter))

	if err != nil {

		//panic(err)

		return nil, err

	}

	//设置头部
	for i, v := range header {

		req.Header.Add(i, v)

	}

	resp, err := client.Do(req)

	if err != nil {

		//panic(err)

		//resp.

		return resp, err

	}

	//defer resp.Body.Close()

	return resp, nil

}

//get获取字符串结果
func GetWithString(url string) (string, error) {

	resp, err := Query(url, "GET", "", nil)

	if err != nil {

		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		//panic(err)

		return "", err

	}

	return string(body), nil

}

/**
header["Accept-Encoding"]="gzip, deflate, br"
*/
func GetSetHeaderWithString(url string, header map[string]string) (string, error) {

	resp, err := Query(url, "GET", "", header)

	if err != nil {

		return "", err
	}

	defer resp.Body.Close()

	r := resp.Body

	if resp.Header.Get("Content-Encoding") == "gzip" {

		r, _ = gzip.NewReader(resp.Body)

	}

	body, err := ioutil.ReadAll(r)

	if err != nil {

		//panic(err)

		return "", err

	}

	return string(body), nil

}

//注意要手动关闭body
func GetWithBody(url string) (io.ReadCloser, error) {

	resp, err := Query(url, "GET", "", nil)

	if err != nil {

		return nil, err
	}

	return resp.Body, nil

}

func DownloadFile(url string, path string) {

	f, err := os.Create(path + ".temp")

	if err != nil {

		f.Close()

		panic(err)

	}

	defer f.Close()

	body, _ := GetWithBody(url)

	defer body.Close()

	_, err = io.Copy(f, body)

	if err != nil {

		f.Close()

		panic(err)

	}

	f.Close()

	if err = os.Rename(path+".temp", path); err != nil {

		panic(err)

	}

}

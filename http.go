package tools

import (
	"compress/gzip"
	"errors"
	_ "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"net/http"
	url_ "net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type HttpSetting struct {
	TimeOut      int                    //超时时间
	Header       map[string]string      //header
	Parameter    map[string]interface{} //参数
	ProxyAddress string                 //代理地址
}

// Query 请求底层函数
func Query(url string, method string, setting HttpSetting) (*http.Response, error) {

	client := http.Client{}

	if setting.TimeOut == 0 {

		setting.TimeOut = 15
	}

	client.Timeout = time.Duration(setting.TimeOut) * time.Second

	//if setting.ProxyAddress != "" {

	netTransport := &http.Transport{
		Proxy: func(r *http.Request) (*url_.URL, error) {

			if setting.ProxyAddress != "" {

				return url_.Parse(setting.ProxyAddress)

			}

			return nil, nil
		},
		DisableKeepAlives:   true,
		MaxIdleConns:        1000,
		MaxIdleConnsPerHost: -1,
		MaxConnsPerHost:     0,
		IdleConnTimeout:     0,
		DisableCompression:  true,
	}

	client.Transport = netTransport

	//}

	var req *http.Request
	var err error

	if method == "GET" {

		p := setting.Parameter

		if len(p) > 0 {

			url += "?"

			for i, v := range p {

				switch key := v.(type) {

				case string:

					url += i + "=" + key + "&"

				case int:

					url += i + "=" + strconv.Itoa(key) + "&"

				case []string:

					for _, vv := range key {

						url += i + "[]=" + vv + "&"

					}

				}

			}

		}

		req, err = http.NewRequest(method, url, nil)

	} else if method == "POST" {

		postForm := ""

		p := setting.Parameter

		//url+="?"

		for i, v := range p {

			switch key := v.(type) {

			case string:

				postForm += i + "=" + key + "&"

			case int:

				postForm += i + "=" + strconv.Itoa(key) + "&"

			case []string:

				for _, vv := range key {

					postForm += i + "[]=" + vv + "&"

				}

			}

		}

		req, err = http.NewRequest(method, url, strings.NewReader(postForm))

		if req != nil {

			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}

	}

	if err != nil {

		//panic(err)

		return nil, err

	}

	//设置头部
	for i, v := range setting.Header {

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

// GetToString get获取字符串结果
func GetToString(url string, setting HttpSetting) (string, error) {

	resp, err := Query(url, "GET", setting)

	if err != nil {

		return "", err
	}

	defer resp.Body.Close()

	read, err := dealBody(resp)

	if err != nil {

		return "", err
	}

	body, err := ioutil.ReadAll(read)

	if err != nil {

		return "", err

	}

	return string(body), nil

}

// GetToStringWithHeader  get获取字符串结果并返回头部信息
func GetToStringWithHeader(url string, setting HttpSetting) (string, http.Header, error) {

	resp, err := Query(url, "GET", setting)

	if err != nil {

		return "", nil, err
	}

	defer resp.Body.Close()

	//resp.Header.Values()

	read, err := dealBody(resp)

	if err != nil {

		return "", nil, err
	}

	body, err := ioutil.ReadAll(read)

	if err != nil {

		return "", nil, err

	}

	return string(body), resp.Header, nil

}

// PostToString post获取字符串结果
func PostToString(url string, setting HttpSetting) (string, error) {

	resp, err := Query(url, "POST", setting)

	if err != nil {

		return "", err
	}

	defer resp.Body.Close()

	read, err := dealBody(resp)

	if err != nil {

		return "", err
	}

	body, err := ioutil.ReadAll(read)

	if err != nil {

		return "", err

	}

	return string(body), nil

}

// GetToBody 注意要手动关闭body
func GetToBody(url string, setting HttpSetting) (io.ReadCloser, error) {

	resp, err := Query(url, "GET", setting)

	if err != nil {

		return nil, err
	}

	read, err := dealBody(resp)

	if err != nil {

		return read, err
	}

	//body, err := ioutil.ReadAll(read)

	return read, nil

}

// GetToResp 注意要手动关闭body
func GetToResp(url string, setting HttpSetting) (*http.Response, error) {

	resp, err := Query(url, "GET", setting)

	if err != nil {

		return nil, err
	}

	return resp, nil

}

// DownloadImage 图片下载
func DownloadImage(url string, path string, setting HttpSetting) error {

	f, err := os.Create(path + ".temp")

	if err != nil {

		//f.Close()

		return err
	}

	defer func() {

		f.Close()

		//删除临时文件
		DeleteFile(path + ".temp")

	}()

	resp, err := GetToResp(url, setting)

	if err != nil {

		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {

		return errors.New("图片下载404")

	}

	contentType := resp.Header.Get("Content-Type")

	//panic(contentType)

	if !(contentType == "image/jpeg" || contentType == "image/png" || contentType == "image/jpg" || contentType == "image/gif") {

		//panic("图片类型错误")

		return errors.New("图片类型错误")

	}

	_, err = io.Copy(f, resp.Body)

	if err != nil {

		f.Close()

		//panic(err)

		return err

	}

	//释放文件占用
	f.Close()

	if err = os.Rename(path+".temp", path); err != nil {

		//panic(err)

		return err
	}

	return nil

}

// DownloadFile 下载文件
func DownloadFile(url string, path string, setting HttpSetting) error {

	f, err := os.Create(path + ".temp")

	if err != nil {

		//f.Close()

		return err
	}

	defer func() {

		f.Close()

		//删除临时文件
		DeleteFile(path + ".temp")

	}()

	resp, err := GetToResp(url, setting)

	if err != nil {

		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 404 {

		return errors.New("404")

	}

	//contentType := resp.Header.Get("Content-Type")
	//
	////panic(contentType)
	//
	//if !(contentType == "image/jpeg" || contentType == "image/png" || contentType == "image/jpg" || contentType == "image/gif") {
	//
	//	//panic("图片类型错误")
	//
	//	return errors.New("图片类型错误")
	//
	//}

	_, err = io.Copy(f, resp.Body)

	if err != nil {

		f.Close()

		//panic(err)

		return err

	}

	//释放文件占用
	f.Close()

	if err = os.Rename(path+".temp", path); err != nil {

		//panic(err)

		return err
	}

	return nil

}

func dealBody(resp *http.Response) (io.ReadCloser, error) {

	r := resp.Body

	var err error

	if resp.Header.Get("Content-Encoding") == "gzip" {

		r, err = gzip.NewReader(resp.Body)

		return r, err

	}

	return r, nil
}

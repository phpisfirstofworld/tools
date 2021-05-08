package http

import (
	"compress/gzip"
	"errors"
	"github.com/PeterYangs/tools"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	url_ "net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// C 客户端结构体
type C struct {
	client    *http.Client
	Transport *http.Transport
}

// R request构造体
type R struct {
	Request   *http.Request
	Parameter map[string]interface{} //参数
	client    *http.Client
	Header    map[string]string //header

}

//client->request->do

// Client 获取客户端
func Client() *C {

	client := http.Client{}

	netTransport := &http.Transport{

		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConnsPerHost:   100,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	client.Transport = netTransport

	c := &C{}

	c.Transport = netTransport

	c.client = &client

	return c

}

// SetTimeout 设置超时时间
func (c *C) SetTimeout(time time.Duration) *C {

	c.client.Timeout = time

	return c

}

// SetProxyAddress 设置代理地址
func (c *C) SetProxyAddress(address string) *C {

	c.Transport.Proxy = func(request *http.Request) (*url_.URL, error) {

		return url_.Parse(address)
	}

	return c

}

func (c *C) Request() *R {

	return &R{client: c.client, Request: &http.Request{}}
}

// SetHeader 设置header
func (r *R) SetHeader(header map[string]string) *R {

	r.Header = header

	return r
}

// SetParameter 设置请求参数
func (r *R) SetParameter(p map[string]interface{}) *R {

	r.Parameter = p

	return r
}

//获取Response
func getResponse(r *R, method string, url string) (*http.Response, error) {

	var req = r.Request
	var err error

	if method == "GET" {

		p := r.Parameter

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

		if err != nil {

			return nil, err
		}

	} else if method == "POST" {

		postForm := ""

		p := r.Parameter

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

		if err != nil {

			return nil, err
		}

		if req != nil {

			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}

	}

	//设置头部
	for i, v := range r.Header {

		req.Header.Add(i, v)
	}

	resp, err := r.client.Do(req)

	if err != nil {

		return resp, err

	}

	return resp, nil

}

func (r *R) GetToString(url string) (string, error) {

	resp, err := getResponse(r, "GET", url)

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
func (r *R) GetToStringWithHeader(url string) (string, http.Header, error) {

	resp, err := getResponse(r, "GET", url)

	if err != nil {

		return "", nil, err
	}

	defer resp.Body.Close()

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
func (r *R) PostToString(url string) (string, error) {

	resp, err := getResponse(r, "POST", url)

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

func dealBody(resp *http.Response) (io.ReadCloser, error) {

	r := resp.Body

	var err error

	if resp.Header.Get("Content-Encoding") == "gzip" {

		r, err = gzip.NewReader(resp.Body)

		return r, err

	}

	return r, nil
}

// GetToBody 注意要手动关闭body
func (r *R) GetToBody(url string) (io.ReadCloser, error) {

	resp, err := getResponse(r, "POST", url)

	if err != nil {

		return nil, err
	}

	read, err := dealBody(resp)

	if err != nil {

		return read, err
	}

	return read, nil

}

// GetToResp 注意要手动关闭body
func (r *R) GetToResp(url string) (*http.Response, error) {

	resp, err := getResponse(r, "GET", url)

	if err != nil {

		return nil, err
	}

	return resp, nil

}

// DownloadImage 图片下载
func (r *R) DownloadImage(url string, path string) error {

	f, err := os.Create(path + ".temp")

	if err != nil {

		//f.Close()

		return err
	}

	defer func() {

		f.Close()

		//删除临时文件
		tools.DeleteFile(path + ".temp")

	}()

	resp, err := r.GetToResp(url)

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
func (r *R) DownloadFile(url string, path string) error {

	f, err := os.Create(path + ".temp")

	if err != nil {

		//f.Close()

		return err
	}

	defer func() {

		f.Close()

		//删除临时文件
		tools.DeleteFile(path + ".temp")

	}()

	resp, err := r.GetToResp(url)

	if err != nil {

		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {

		return errors.New("http code : " + strconv.Itoa(resp.StatusCode) + ",link:" + url)

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

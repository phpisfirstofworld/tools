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
	Header    map[string]string //全局头部
}

// R request构造体
type R struct {
	Request    *http.Request
	Parameter  map[string]interface{} //参数
	c          *C
	Header     map[string]string //header
	ReTryTimes int               //重试次数
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

	//默认超时时间
	c.client.Timeout = 15 * time.Second

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

// SetHeader 设置全局header
func (c *C) SetHeader(header map[string]string) *C {

	c.Header = header

	return c
}

func (c *C) Request() *R {

	return &R{c: c, Request: &http.Request{}, ReTryTimes: 0}
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

// SetReTryTimes 设置重试次数
func (r *R) SetReTryTimes(times int) *R {

	if times < 0 {

		times = 0
	}

	r.ReTryTimes = times

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

			url = resolveInterface(p, url, []string{})

		}

		if tools.SubStr(url, len(url)-1, -1) == "&" {

			url = tools.SubStr(url, 0, len(url)-1)
		}

		req, err = http.NewRequest(method, url, nil)

		if err != nil {

			return nil, err
		}

	} else if method == "POST" {

		postForm := ""

		p := r.Parameter

		postForm = resolveInterface(p, postForm, []string{})

		if tools.SubStr(postForm, len(postForm)-1, -1) == "&" {

			postForm = tools.SubStr(postForm, 0, len(postForm)-1)
		}

		req, err = http.NewRequest(method, url, strings.NewReader(postForm))

		if err != nil {

			return nil, err
		}

		if req != nil {

			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}

	}

	var finalHeader = make(map[string]string)

	//设置全局头部
	for i, v := range r.c.Header {

		finalHeader[i] = v
	}

	//设置request头部,request中设置的header会覆盖全局header
	for i, v := range r.Header {

		finalHeader[i] = v
	}

	for s, s2 := range finalHeader {

		req.Header.Add(s, s2)
	}

	var resp *http.Response
	var e error

	for i := 0; i < r.ReTryTimes+1; i++ {

		resp, e = r.c.client.Do(req)

		if e != nil {

			continue
		}

		break

	}

	if e != nil {

		return resp, e
	}

	if resp.StatusCode != 200 {

		return resp, errors.New("status is " + strconv.Itoa(resp.StatusCode))
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

//解析参数拼接参数字符串
func resolveInterface(p map[string]interface{}, form string, parentName []string) string {

	if len(p) > 0 {

		for i, v := range p {

			switch key := v.(type) {

			case string:

				form += getKey(parentName, i) + "=" + key + "&"

			case int:

				form += getKey(parentName, i) + "=" + strconv.Itoa(key) + "&"

			case []string:

				for _, vv := range key {

					form += getKey(parentName, i) + "[]=" + vv + "&"

				}

			case []int:

				for _, vv := range key {

					form += getKey(parentName, i) + "[]=" + strconv.Itoa(vv) + "&"

				}

			case map[string]interface{}:

				t := append(parentName, i)

				form = resolveInterface(key, form, t)

			}

		}

	}

	return form
}

func getKey(parentName []string, ii string) string {

	f := ""

	for i, s := range parentName {

		if i == 0 {

			f += s

		} else {

			f += "[" + s + "]"

		}

	}

	if len(parentName) > 0 {

		f += "[" + ii + "]"

	} else {

		f += ii

	}

	return f

}

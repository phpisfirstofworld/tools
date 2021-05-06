package http

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	url_ "net/url"
	"strconv"
	"strings"
	"time"
)

type HttpClient struct {
	client      *http.Client
	httpSetting HttpSetting
}

type HttpSetting struct {
	TimeOut      int                    //超时时间
	Header       map[string]string      //header
	Parameter    map[string]interface{} //参数
	ProxyAddress string                 //代理地址
}

func Client(setting HttpSetting) *HttpClient {

	//http.DefaultTransport.(http.Transport).MaxIdleConnsPerHost = 1000
	//http.DefaultTransport.(http.Transport).MaxIdleConns = 1000

	//http.DefaultTransport.(http.Transport).MaxIdleConnsPerHost=1000

	client := http.Client{}

	if setting.TimeOut == 0 {

		setting.TimeOut = 15
	}

	client.Timeout = time.Duration(setting.TimeOut) * time.Second

	//if setting.ProxyAddress != "" {

	netTransport := &http.Transport{
		//Proxy: func(r *http.Request) (*url_.URL, error) {
		//
		//	if setting.ProxyAddress != "" {
		//
		//		return url_.Parse(setting.ProxyAddress)
		//
		//	}
		//
		//	return nil, nil
		//},
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

	if setting.ProxyAddress != "" {

		netTransport.Proxy = func(request *http.Request) (*url_.URL, error) {

			return url_.Parse(setting.ProxyAddress)
		}

	}

	//defaultRoundTripper := http.DefaultTransport
	//defaultTransportPointer, ok := defaultRoundTripper.(*http.Transport)
	//if !ok {
	//	panic(fmt.Sprintf("defaultRoundTripper not an *http.Transport"))
	//}
	//defaultTransport := *defaultTransportPointer // dereference it to get a copy of the struct that the pointer points to
	//defaultTransport.MaxIdleConns = 100
	//defaultTransport.MaxIdleConnsPerHost = 100
	////defaultTransport.Proxy
	//
	//client.Transport = &defaultTransport
	client.Transport = netTransport

	h := HttpClient{}

	h.client = &client

	return &h

}

func getResponse(h HttpClient, method string, url string) (*http.Response, error) {

	var req *http.Request
	var err error

	setting := h.httpSetting

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

		if err != nil {

			return nil, err
		}

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

		if err != nil {

			return nil, err
		}

		if req != nil {

			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}

	}

	//设置头部
	for i, v := range setting.Header {

		req.Header.Add(i, v)

	}

	resp, err := h.client.Do(req)

	if err != nil {

		return resp, err

	}

	return resp, nil

}

func (h HttpClient) GetToString(url string) (string, error) {

	//resp, err := Query(url, "GET", setting)

	resp, err := getResponse(h, "GET", url)

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

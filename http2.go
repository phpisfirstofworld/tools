package tools

import (
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	client http.Client
}

func Client(setting HttpSetting) HttpClient {

	h := HttpClient{}

	client := http.Client{}

	h.client = client

	return h

}

func (h HttpClient) Get(url string) (string, error) {

	rep, err := h.client.Get(url)

	if err != nil {

		return "", err
	}

	defer rep.Body.Close()

	body, err := ioutil.ReadAll(rep.Body)

	return string(body), nil
}

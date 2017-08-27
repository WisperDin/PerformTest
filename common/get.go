package common

import (
	"../conf"
	"io/ioutil"
	"net/http"
)

//参数为url，返回回应的body
func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(conf.App.ServerUrl + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

package common

import (
	"io/ioutil"
	"net/http"
	"../conf"
)

//参数为url，返回回应的body
func HttpGet(url string) (string, error) {
	resp, err := http.Get(conf.App.ServerUrl+url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

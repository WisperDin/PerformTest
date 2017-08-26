package common

import (
	"strings"
	"io/ioutil"
	"net/http"
	"../conf"
)

//参数为url，参数
func HttpPost(url,param string) (string, error) {
	resp, err := http.Post(conf.App.ServerUrl+url,
		"application/x-www-form-urlencoded",
		strings.NewReader(param))
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

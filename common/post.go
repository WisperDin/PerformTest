package common

import (
	"../conf"
	"io/ioutil"
	"net/http"
	"strings"
)

//参数为url，参数
func HttpPost(url, param string) ([]byte, error) {
	resp, err := http.Post(conf.App.ServerUrl+url,
		"application/x-www-form-urlencoded",
		strings.NewReader(param))
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

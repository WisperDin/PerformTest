package module

import (
	"../common"
	"encoding/json"
	"fmt"
	"log"
)

func Login(account, pwd, roleType, osType, oId string) *R {

	param := fmt.Sprintf("account=%s&pwd=%s&osType=%s&roleType=%s&oId=%s", account, pwd, osType, roleType, oId)
	body, err := common.HttpPost("/login", param)
	if err != nil {
		log.Println(err)
	}
	var fb R
	err = json.Unmarshal(body, &fb)
	if err != nil {
		log.Println(err)
	}
	return &fb
}

func Logout(account, roleType, osType, token string) *R {
	param := fmt.Sprintf("account=%s&osType=%s&roleType=%s&token=%s", account, osType, roleType, token)
	body, err := common.HttpGet("/logout?" + param)
	if err != nil {
		log.Println(err)
	}
	var fb R
	err = json.Unmarshal(body, &fb)
	if err != nil {
		log.Println(err)
	}
	return &fb
}

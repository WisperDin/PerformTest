package module

import (
	"../common"
	"fmt"
	"log"
)

func Login(account,pwd,roleType,osType string){
	param:=fmt.Sprintf("account=%s&pwd=%s&osType=%s&roleType=%s",account,pwd,osType,roleType)
	body, err := common.HttpPost("/login",param)
	if err != nil {
		log.Println(err)
	}
	log.Println(body)
}

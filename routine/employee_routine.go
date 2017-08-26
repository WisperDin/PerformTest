package routine

import (
	"../module"
	"log"
	"fmt"
)
//医护人员 	加载审核通过的机构列表->登录->注销
const (
	CORRECT = 1
)

const SUCCESS_CODE  = 200

func EleRoutine(mode int,rid int){
	fb1:=module.ReadOrgList(1)
	//oidRaw:=fb1.Data.(map[string]interface{})["organizations"].([]interface{})[0].(map[string]interface{})["id"].(float64)
	//oid:=strconv.FormatFloat(oidRaw,'f',0,64)
	if mode == CORRECT{
		if fb1.Code!=SUCCESS_CODE{
			log.Fatalln(fmt.Sprintf("fail in %d",rid))
		}
	}

/*	fb2:=module.Login("100000000","123456","employee","web",oid)
	if mode == CORRECT{
		if fb1.Code!=SUCCESS_CODE{
			log.Println(fmt.Sprintf("fail in %d",rid))
		}
	}
	//log.Println(fb2)
	userToken := fb2.Data.(map[string]interface{})["token"].(string)

	module.Logout("100000000","employee","web",userToken)
	if mode == CORRECT{
		if fb1.Code!=SUCCESS_CODE{
			log.Println(fmt.Sprintf("fail in %d",rid))
		}
	}*/
	//log.Println(fb3)

}

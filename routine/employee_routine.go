package routine

import (
	"../module"
	"log"
	"fmt"
	"strconv"
)
//医护人员 	加载审核通过的机构列表->登录->注销
const (
	CORRECT = 1
)

const SUCCESS_CODE  = 200

//找到对应名字的机构id
func FindOrgId(fb *module.R)(oid string){
	var oidRaw float64
	oidRaw=0
	for _,item:=range fb.Data.(map[string]interface{})["organizations"].([]interface{}){
		if item.(map[string]interface{})["name"]=="123有限公司"{
			oidRaw=item.(map[string]interface{})["id"].(float64)
		}
	}
	if oidRaw == 0{
		return ""
	}
	oid=strconv.FormatFloat(oidRaw,'f',0,64)
	return
}

func EleRoutine(mode int,rid int){
	fb1:=module.ReadOrgList(1)
	if mode == CORRECT{
		if fb1.Code!=SUCCESS_CODE{
			log.Fatalln(fmt.Sprintf("fail in %d",rid))
		}
	}
	//找到对应名字的机构id
	oid:=FindOrgId(fb1)
	if oid==""{
		log.Fatalln(fmt.Sprintf("fail in %d",rid))
	}

	fb2:=module.Login("100000000","123456","employee","web",oid)
	if mode == CORRECT{
		if fb1.Code!=SUCCESS_CODE{
			log.Println(fmt.Sprintf("fail in %d",rid))
		}
	}
	userToken := fb2.Data.(map[string]interface{})["token"].(string)

	module.Logout("100000000","employee","web",userToken)
	if mode == CORRECT{
		if fb1.Code!=SUCCESS_CODE{
			log.Println(fmt.Sprintf("fail in %d",rid))
		}
	}


}

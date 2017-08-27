package routine

import (
	."../module"
	"fmt"
	"log"
	"strconv"
	"../report"
)

//医护人员 	加载审核通过的机构列表->登录->注销

//todo....数据自动准备,往数据库插入测试用登录的机构,测试用员工账号()
func elePrepare() {

}

//找到对应名字的机构id
func FindOrgId(fb *R) (oid string) {
	var oidRaw float64
	oidRaw = 0
	for _, item := range fb.Data.(map[string]interface{})["organizations"].([]interface{}) {
		if item.(map[string]interface{})["name"] == "123有限公司" {
			oidRaw = item.(map[string]interface{})["id"].(float64)
		}
	}
	if oidRaw == 0 {
		return ""
	}
	oid = strconv.FormatFloat(oidRaw, 'f', 0, 64)
	return
}

func EleRoutine(mode int, rid int) {
	routineRes:=report.RoutineResult{}
	routineRes.IsFinish=false
	fb1 := ReadOrgList(1)
	if mode == CORRECT {
		if fb1.Code != SUCCESS_CODE {
			routineRes.FailPos=0
			log.Fatalln(fmt.Sprintf("fail in rid=%d pos=ReadOrgList", rid))
		}
	}
	//找到对应名字的机构id
	oid := FindOrgId(fb1)
	if oid == "" {
		log.Fatalln(fmt.Sprintf("fail in %d ", rid))
	}

	fb2 := Login("100000000", "123456", "employee", "web", oid)
	if mode == CORRECT {
		if fb1.Code != SUCCESS_CODE {
			routineRes.FailPos=1
			log.Fatalln(fmt.Sprintf("fail in rid=%d pos=Login", rid))
		}
	}
	userToken := fb2.Data.(map[string]interface{})["token"].(string)

	Logout("100000000", "employee", "web", userToken)
	if mode == CORRECT {
		if fb1.Code != SUCCESS_CODE {
			routineRes.FailPos=2
			log.Fatalln(fmt.Sprintf("fail in rid=%d pos=Logout", rid))
		}
	}
	routineRes.IsFinish=true
	report.RoutineChan<-routineRes
}

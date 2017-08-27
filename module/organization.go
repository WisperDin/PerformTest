package module

import (
	"../common"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

/**
机构相关
*/

/*
* @param 机构状态 0表示待审核,1表示通过审核,2表示被拒绝
 */
func ReadOrgList(auditState int) *R {
	param := fmt.Sprintf("auditState=%s", strconv.Itoa(auditState))
	body, err := common.HttpGet("/organizations/read?" + param)
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

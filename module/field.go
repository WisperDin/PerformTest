package module

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type RE struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

const (
	CORRECT = 1
)

const SUCCESS_CODE = 200

/*//检查服务器反馈是否成功
func (this* Reporter)chkFb(fb *R){
	if fb.Code==SUCCESS_CODE{
		this.SuccessUserNum++
	}else{
		this.FailUserNum++
	}
}*/

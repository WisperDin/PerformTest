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



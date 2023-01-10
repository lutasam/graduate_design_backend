package bo

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PingResponse struct {
	Pong string `json:"pong"`
}

type HelloResponse struct {
	Hello string `json:"hello"`
}

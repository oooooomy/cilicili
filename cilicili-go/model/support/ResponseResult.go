package support

type ResponseResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success() *ResponseResult {
	return &ResponseResult{
		Code: 200,
		Msg:  "request ok",
		Data: nil,
	}
}

func SuccessWithData(data interface{}) *ResponseResult {
	return &ResponseResult{
		Code: 200,
		Msg:  "request ok",
		Data: data,
	}
}

func Error(code int, msg string) *ResponseResult {
	return &ResponseResult{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

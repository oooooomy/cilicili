package email

//验证码过期时间
const expTime = 1000 * 60 * 10

var codeMap = make(map[string]*emailCode)

type emailCode struct {
	value string
	exp   int64
}

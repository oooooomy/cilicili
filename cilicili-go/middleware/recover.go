package middleware

import (
	"cilicili-go/model/support"
	"github.com/kataras/iris/v12"
	"runtime/debug"
)

// Recover 中间件
func Recover(c iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			_, _ = c.JSON(support.Error(500, errorToString(err)))
			return
		}
	}()
	c.Next()
}

// recover错误，转string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

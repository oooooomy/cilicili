package email

import (
	"cilicili-go/middleware"
	"github.com/kataras/iris/v12"
)

func RegisterAllRouter(app *iris.Application) {

	//使用全局中间件
	app.Use(middleware.Recover)

	//获取邮箱验证码
	app.Get("/code", getEmailCode)

	//发送邮箱
	app.Post("/simple", sendSimpleEmail)

	//检验邮箱
	app.Get("/check", checkEmailCode)

}

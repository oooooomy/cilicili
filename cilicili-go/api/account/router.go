package account

import (
	"cilicili-go/middleware"
	"github.com/kataras/iris/v12"
)

func RegisterAllRouter(app *iris.Application) {

	//使用全局中间件
	app.Use(middleware.Recover)

	//小程序用户登录
	app.Post("/user/login", userLogin)

	//更新用户信息
	app.Put("/user/update", userUpdate)

	//根据id获取用户信息
	app.Get("/user/{id}", getUserInfo)

	//获取所有管理员
	app.Get("/admin", findAllAdmin)

	//管理员密码登录
	app.Post("/admin/loginByPassword", adminLoginByPassword)

	//管理员验证码登录
	app.Get("/admin/loginByEmail", adminLoginByEmail)

	//添加管理员
	app.Post("/admin/create", createAdmin)

	//删除管理员
	app.Delete("/admin/{id}", deleteAdmin)

	//添加关注
	app.Post("/follow/create", createFollow)

	//取消关注
	app.Delete("/follow/from/{from}/to/{to}", deleteFollow)

	//获取某个用户的关注列表
	app.Get("/follow/from/{id}", getUserFollowList)

	//查看是否关注某个用户
	app.Get("/follow/from/{from}/to/{to}", isFollow)

}

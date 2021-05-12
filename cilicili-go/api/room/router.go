package room

import (
	"cilicili-go/middleware"
	"github.com/kataras/iris/v12"
)

func RegisterAllRouter(app *iris.Application) {

	//使用全局中间件
	app.Use(middleware.Recover)

	app.Get("/ws/conn", wsHandler)

	app.Get("/findAll", findAllRoom)

	app.Get("/findById/{id}", findById)

	//开播
	app.Get("/up/{id}", up)

	//下播
	app.Get("/down/{id}", down)

	app.Get("/findByStatus/{status}", findByStatus)

	app.Get("/findByLiving/{living}", findByLiving)

	app.Post("/create", saveRoom)

}

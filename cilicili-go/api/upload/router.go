package upload

import "github.com/kataras/iris/v12"

func RegisterAllRouter(app *iris.Application) {

	app.Post("/{type}", postFile)

}

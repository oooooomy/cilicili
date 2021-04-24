package video

import (
	"cilicili-go/middleware"
	"github.com/kataras/iris/v12"
)

func RegisterAllRouter(app *iris.Application) {

	app.Use(middleware.Recover)

	app.Get("/", findAllVideo)

	app.Get("/findByUser/{userId}", findVideoByUser)

	app.Get("/findByMusic/{musicId}", findVideoByMusic)

	app.Post("/", saveVideo)

	app.Put("/", updateVideo)

	app.Delete("/{id}", deleteVideo)

	app.Post("/upload/mp4", uploadMp4)

	app.Post("/upload/mp3", uploadMp3)

}

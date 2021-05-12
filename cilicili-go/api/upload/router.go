package upload

import (
	"cilicili-go/conf"
	"cilicili-go/middleware"
	"github.com/kataras/iris/v12"
)

func RegisterAllRouter(app *iris.Application) {

	//使用全局中间件
	app.Use(middleware.Recover)

	//获取静态文件
	app.HandleDir("/file", conf.Config.Upload.Path)

	//上传文件
	app.Post("/file", putFile)

	/*//获取文件
	app.Get("/file/{name}", getFile)*/

	//获取所有音乐
	app.Get("/music", findAllMusic)

	//添加音乐
	app.Post("/music", saveMusic)

	//删除音乐
	app.Delete("/music/{id}", deleteMusic)

	//添加评论
	app.Post("/comment", saveComment)

	//获取video的所有评论
	app.Get("/comment/video/{id}", findByVideo)

	app.Post("/video", saveVideo)

	app.Put("/video/{id}/setStatus", setVideoStatus)

	app.Get("/video", findAllVideo)

	app.Post("/video/like", addVideoLike)

	app.Get("/video/findByMusicId", findByMusicId)

	app.Get("/video/findAllByFollow", findAllByFollow)

	app.Get("/video/findByUserId", findByUserId)

	app.Get("/video/findByStatus", findByStatus)

	app.Get("/video/findUserLikeAndUpload/{id}", findUserLikeAndUpload)

}

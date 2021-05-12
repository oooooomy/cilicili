package main

import (
	"cilicili-go/api/upload"
	"cilicili-go/client"
	"cilicili-go/conf"
	_ "cilicili-go/conf"
	db "cilicili-go/db/upload"
	"github.com/kataras/iris/v12"
	"os"
	"strconv"
)

func main() {

	//获取启动端口
	port := os.Args[len(os.Args)-1]

	i, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	//注册到eureka注册中心
	client.RegisterToEureka(conf.Config.Eureka.Address, i, "upload-service")

	//连接数据库
	db.ConnVideoDatabase()

	app := iris.New()
	upload.RegisterAllRouter(app)
	_ = app.Run(iris.Addr(":" + port))

}

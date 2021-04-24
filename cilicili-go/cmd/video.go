package main

import (
	"cilicili-go/api/video"
	"cilicili-go/conf"
	_ "cilicili-go/conf"
	db "cilicili-go/db/video"
	_ "cilicili-go/service/video"
	"github.com/kataras/iris/v12"
	eureka "github.com/xuanbo/eureka-client"
	"os"
	"strconv"
)

var videoServerPort string

func main() {
	//连接数据库
	db.ConnVideoDatabase()
	app := iris.New()
	video.RegisterAllRouter(app)
	_ = app.Run(iris.Addr(":" + videoServerPort))
}

func init() {

	//获取启动端口
	videoServerPort = os.Args[len(os.Args)-1]

	i, err := strconv.Atoi(videoServerPort)
	if err != nil {
		panic(err)
	}

	//注册到eureka注册中心
	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           conf.Config.Eureka.Address,
		App:                   "video-service",
		Port:                  i,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
	})
	client.Start()

}

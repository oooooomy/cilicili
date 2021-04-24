package main

import (
	"cilicili-go/api/video"
	"cilicili-go/conf"
	_ "cilicili-go/conf"
	_ "cilicili-go/handler/video"
	"github.com/kataras/iris/v12"
	eureka "github.com/xuanbo/eureka-client"
	"os"
	"strconv"
)

var userServerPort string

func main() {
	app := iris.New()
	video.RegisterAllRouter(app)
	err := app.Run(iris.Addr(":" + userServerPort))
	if err != nil {
		panic(err)
	}
}

func init() {

	//获取启动端口
	userServerPort = os.Args[len(os.Args)-1]

	i, err := strconv.Atoi(userServerPort)
	if err != nil {
		panic(err)
	}

	//注册到eureka注册中心
	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           conf.Config.Eureka.Address,
		App:                   "user-service",
		Port:                  i,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
	})
	client.Start()

}

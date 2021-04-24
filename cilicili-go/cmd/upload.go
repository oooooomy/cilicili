package main

import (
	"cilicili-go/api/upload"
	"cilicili-go/conf"
	_ "cilicili-go/conf"
	"cilicili-go/service"
	"github.com/kataras/iris/v12"
	eureka "github.com/xuanbo/eureka-client"
	"os"
	"strconv"
)

var uploadServerPort string

func main() {
	//加载阿里云oss客户端
	service.LoadOssClient()
	app := iris.New()
	upload.RegisterAllRouter(app)
	_ = app.Run(iris.Addr(":" + uploadServerPort))
}

func init() {

	//获取启动端口
	uploadServerPort = os.Args[len(os.Args)-1]

	i, err := strconv.Atoi(uploadServerPort)
	if err != nil {
		panic(err)
	}

	//注册到eureka注册中心
	client := eureka.NewClient(&eureka.Config{
		DefaultZone:           conf.Config.Eureka.Address,
		App:                   "upload-service",
		Port:                  i,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
	})
	client.Start()

}

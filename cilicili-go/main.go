package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	app.HandleDir("/file","/Users/soanr/Desktop/temp")
	err := app.Run(iris.Addr(":10000"))
	if err != nil {
		panic(err)
	}
}
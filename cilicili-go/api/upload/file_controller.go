package upload

import (
	"cilicili-go/conf"
	"cilicili-go/model/support"
	"github.com/go-basic/uuid"
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"strings"
)

//此方法获取mp4时小程序开发工具正常
//真机不显示
func getFile(c iris.Context) {
	name := c.Params().Get("name")
	if name == "" {
		_, _ = c.JSON(support.Error(400, "file name missing"))
		return
	}
	file, err := os.Open(conf.Config.Upload.Path + name)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	defer func() {
		_ = file.Close()
	}()
	c.Header("Server", "cilicili")
	c.Header("Connection","keep-alive")
	c.Header("Content-Type", "video/mp4")
	_, err = io.Copy(c.ResponseWriter(), file)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
}

func putFile(c iris.Context) {
	file, header, err := c.FormFile("file")
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	split := strings.Split(header.Filename, ".")
	fileType := split[len(split)-1]
	fileName := uuid.New() + "." + fileType
	f, err := os.OpenFile(conf.Config.Upload.Path+fileName, os.O_WRONLY|os.O_CREATE, 0766)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	defer func() {
		_ = file.Close()
	}()
	_, err = io.Copy(f, file)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.SuccessWithData(fileName))
}

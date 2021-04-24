package upload

import (
	"cilicili-go/model/support"
	"cilicili-go/service"
	"github.com/kataras/iris/v12"
	"strings"
)

var uploadService = service.UploadServiceImpl{}

func postFile(c iris.Context) {
	fileType := c.Params().Get("type")
	file, header, err := c.FormFile("file")
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	//分割字符
	split := strings.Split(header.Filename, ".")
	//防止文件名包含 "."
	if split[len(split)-1] != fileType {
		_, _ = c.JSON(support.Error(400, "Disallowed file type"))
		return
	}
	url, err := uploadService.Upload(file, fileType)
	if err != nil {
		_, _ = c.JSON(support.Error(500, err.Error()))
		return
	}
	_, _ = c.JSON(support.SuccessWithData(url))
}

package upload

import (
	"cilicili-go/model/entity"
	"cilicili-go/model/support"
	"cilicili-go/service/upload"
	"github.com/kataras/iris/v12"
)

var commentService = new(upload.CommentService)

func saveComment(c iris.Context) {
	comment := new(entity.Comment)
	err := c.ReadJSON(comment)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	commentService.Create(comment)
	_, _ = c.JSON(support.Success())
}

func findByVideo(c iris.Context) {
	_, _ = c.JSON(support.SuccessWithData(commentService.FindAllByVideoId(c.Params().Get("id"))))
}

package upload

import (
	"cilicili-go/model/entity"
	"cilicili-go/model/support"
	"cilicili-go/service/upload"
	"github.com/kataras/iris/v12"
)

var musicService = new(upload.MusicService)

func saveMusic(c iris.Context) {
	m := &entity.Music{}
	err := c.ReadJSON(m)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	musicService.Create(m)
	_, _ = c.JSON(support.Success())
}

func deleteMusic(c iris.Context) {
	id := c.Params().Get("id")
	musicService.Delete(id)
	_, _ = c.JSON(support.Success())
}

func findAllMusic(c iris.Context) {
	_, _ = c.JSON(support.SuccessWithData(musicService.FindAll()))
}

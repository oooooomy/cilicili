package video

import (
	"cilicili-go/model/entity"
	"cilicili-go/model/support"
	"cilicili-go/service"
	"github.com/kataras/iris/v12"
)

var videoService = new(service.VideoServiceImpl)

func findAllVideo(c iris.Context) {
	videos := videoService.FindAllVideo()
	_, _ = c.JSON(support.SuccessWithData(videos))
}

func findVideoByUser(c iris.Context) {
	userId := c.Params().Get("userId")
	videos := videoService.FindVideoByUser(userId)
	_, _ = c.JSON(support.SuccessWithData(videos))
}

func findVideoByMusic(c iris.Context) {
	musicId := c.Params().Get("musicId")
	videos := videoService.FindVideoByMusic(musicId)
	_, _ = c.JSON(support.SuccessWithData(videos))
}

func saveVideo(c iris.Context) {
	video := &entity.Video{}
	err := c.ReadJSON(video)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	video = videoService.CreateVideo(video)
	_, _ = c.JSON(support.SuccessWithData(video))
}

func deleteVideo(c iris.Context) {
	id := c.Params().Get("id")
	videoService.DeleteVideo(id)
	_, _ = c.JSON(support.Success())
}

func updateVideo(c iris.Context) {
	video := &entity.Video{}
	err := c.ReadJSON(video)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.SuccessWithData(videoService.UpdateVideo(video)))
}

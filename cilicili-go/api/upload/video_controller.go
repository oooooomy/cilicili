package upload

import (
	"cilicili-go/model/entity"
	"cilicili-go/model/support"
	"cilicili-go/service/upload"
	"github.com/kataras/iris/v12"
	"strconv"
)

var videoService = new(upload.VideoService)

func saveVideo(c iris.Context) {
	video := new(entity.Video)
	err := c.ReadJSON(video)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	videoService.Create(video)
	_, _ = c.JSON(support.Success())
}

func setVideoStatus(c iris.Context) {
	id := c.Params().Get("id")
	status := c.URLParam("status")
	b, err := strconv.ParseBool(status)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	videoService.SetStatus(id, b)
	_, _ = c.JSON(support.Success())
}

func findAllVideo(c iris.Context) {
	_, _ = c.JSON(support.SuccessWithData(videoService.FindAll()))
}

func addVideoLike(c iris.Context) {
	like := new(entity.VideoLike)
	err := c.ReadJSON(like)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	err = videoService.Like(like)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.Success())
}

func findByMusicId(c iris.Context) {
	arr := videoService.FindAllByUserId(c.URLParam("musicId"))
	_, _ = c.JSON(support.SuccessWithData(arr))
}

func findAllByFollow(c iris.Context) {
	param := c.URLParam("userId")
	videos := videoService.FindAllByFollow(param)
	_, _ = c.JSON(support.SuccessWithData(videos))
}

func findByUserId(c iris.Context) {
	_, _ = c.JSON(support.SuccessWithData(videoService.FindAllByUserId(c.URLParam("userId"))))
}

func findByStatus(c iris.Context) {
	status := c.URLParam("status")
	b, err := strconv.ParseBool(status)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.SuccessWithData(videoService.FindAllByStatus(b)))
}

func findUserLikeAndUpload(c iris.Context) {
	m := make(map[string][]*entity.Video)
	m["uploadList"] = videoService.FindAllByUserId(c.Params().Get("id"))
	m["likeList"] = videoService.FindByUserLike(c.Params().Get("id"))
	_, _ = c.JSON(support.SuccessWithData(m))
}

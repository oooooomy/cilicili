package video

import (
	"cilicili-go/model/entity"
	"cilicili-go/model/support"
	videoService "cilicili-go/service/video"
	"github.com/kataras/iris/v12"
	"strings"
)

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

func uploadMp4(c iris.Context) {
	file, header, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	filename := header.Filename
	//分割字符
	split := strings.Split(filename, ".")
	//防止文件名包含 "."
	if split[len(split)-1] != "mp4" {
		_, _ = c.JSON(support.Error(400, "Disallowed file type"))
		return
	}
	name, err := videoService.PutMp4(file)
	if err != nil {
		panic(err)
	}
	_, _ = c.JSON(support.SuccessWithData(name))
}

func uploadMp3(c iris.Context) {
	file, header, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	filename := header.Filename
	//分割字符
	split := strings.Split(filename, ".")
	//防止文件名包含 "."
	if split[len(split)-1] != "mp3" {
		_, _ = c.JSON(support.Error(400, "Disallowed file type"))
		return
	}
	name, err := videoService.PutMp3(file)
	if err != nil {
		panic(err)
	}
	_, _ = c.JSON(support.SuccessWithData(name))
}

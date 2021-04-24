package video

import (
	"cilicili-go/db/video"
	"cilicili-go/model/entity"
	"io"
	"time"
)

var videoDao = new(video.BaseVideoDao)

func CreateVideo(video *entity.Video) *entity.Video {
	video.CreateAt = time.Now().Unix()
	videoDao.Save(video)
	return video
}

func UpdateVideo(video *entity.Video) *entity.Video {
	videoDao.Update(video)
	return video
}

func FindAllVideo() []*entity.Video {
	return videoDao.FindAll()
}

func FindVideoByUser(id string) []*entity.Video {
	return videoDao.FindByUserId(id)
}

func FindVideoByMusic(id string) []*entity.Video {
	return videoDao.FindByMusicId(id)
}

func DeleteVideo(id string) {
	videoDao.Delete(id)
}

func PutMp4(data io.Reader) (string, error) {
	return putObject("mp4", data)
}

func PutMp3(data io.Reader) (string, error) {
	return putObject("mp3", data)
}

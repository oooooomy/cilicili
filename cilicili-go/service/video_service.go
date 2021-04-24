package service

import (
	"cilicili-go/db/video"
	"cilicili-go/model/entity"
	"time"
)

type VideoService interface {
	CreateVideo(video *entity.Video) *entity.Video
	UpdateVideo(video *entity.Video) *entity.Video
	FindAllVideo() []*entity.Video
	FindVideoByUser(id string) []*entity.Video
	FindVideoByMusic(id string) []*entity.Video
	DeleteVideo(id string)
}

type VideoServiceImpl struct{}

var videoDao = new(video.BaseVideoDao)

func (v *VideoServiceImpl) CreateVideo(video *entity.Video) *entity.Video {
	video.CreateAt = time.Now().Unix()
	videoDao.Save(video)
	return video
}

func (v *VideoServiceImpl) UpdateVideo(video *entity.Video) *entity.Video {
	videoDao.Update(video)
	return video
}

func (v *VideoServiceImpl) FindAllVideo() []*entity.Video {
	return videoDao.FindAll()
}

func (v *VideoServiceImpl) FindVideoByUser(id string) []*entity.Video {
	return videoDao.FindByUserId(id)
}

func (v *VideoServiceImpl) FindVideoByMusic(id string) []*entity.Video {
	return videoDao.FindByMusicId(id)
}

func (v *VideoServiceImpl) DeleteVideo(id string) {
	videoDao.Delete(id)
}

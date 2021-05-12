package upload

import (
	"cilicili-go/client"
	"cilicili-go/conf"
	db "cilicili-go/db/upload"
	"cilicili-go/model/entity"
	"errors"
	"github.com/go-basic/uuid"
	"os"
	"os/exec"
)

type BaseVideoService interface {
	Create(video *entity.Video)
	SetStatus(id string, status bool)
	FindAll() []*entity.Video
	Like(like *entity.VideoLike) error
	FindByUserLike(userId string) []*entity.Video
	FindAllByStatus(status bool) []*entity.Video
	FindAllByMusicId(musicId string) []*entity.Video
	FindAllByFollow(userId string) []*entity.Video
	FindAllByUserId(userId string) []*entity.Video
}

type VideoService struct{}

var (
	videoDao     = new(db.VideoDao)
	videoLikeDao = new(db.VideoLikeDao)
)

func (v *VideoService) Create(video *entity.Video) {
	//mp3合成mp4
	if video.UseMusic {
		oldPath := video.Url
		newPath := uuid.New() + ".mp4"
		args := []string{
			"-i",
			conf.Config.Upload.Path + oldPath,
			"-i",
			conf.Config.Upload.Path + video.MusicUrl,
			"-c:v", "copy", "-shortest", "-c:a", "aac", "-strict", "experimental", "-map", "0:v:0", "-map", "1:a:0",
			conf.Config.Upload.Path + newPath,
		}
		cmd := exec.Command(conf.Config.Upload.Ffmpeg, args...)
		err := cmd.Run()
		if err != nil {
			panic(err)
		}
		defer func() {
			//删除处理前的原视频
			_ = os.Remove(conf.Config.Upload.Path + oldPath)
		}()
		video.Url = newPath
	}
	videoDao.Save(video)
}

func (v *VideoService) SetStatus(id string, status bool) {
	videoDao.SetStatusById(id, status)
}

func (v *VideoService) FindAll() []*entity.Video {
	return videoDao.FindAll()
}

func (v *VideoService) FindAllByStatus(status bool) []*entity.Video {
	return videoDao.FindAllByStatus(status)
}

func (v *VideoService) FindByUserLike(userId string) (videos []*entity.Video) {
	likes := videoLikeDao.FindAllByUserId(userId)
	for _, like := range likes {
		videos = append(videos, videoDao.FindById(like.VideoId))
	}
	return
}

func (v *VideoService) FindAllByMusicId(musicId string) []*entity.Video {
	return videoDao.FindAllByMusicId(musicId)
}

func (v *VideoService) FindAllByFollow(userId string) []*entity.Video {
	c := &client.AccountServiceClient{}
	follows := c.FindUserFollow(userId)
	arr := make([]*entity.Video, 0)
	for _, follow := range follows {
		arr = append(arr, v.FindAllByUserId(follow.ToUserId)...)
	}
	return arr
}

func (v *VideoService) FindAllByUserId(userId string) []*entity.Video {
	return videoDao.FindAllByUserId(userId)
}

func (v *VideoService) Like(like *entity.VideoLike) error {
	if videoLikeDao.ExistsByUserIdAndVideoId(like.UserId, like.VideoId) {
		return errors.New("你已经点过赞了哦")
	}
	videoDao.AddVideoLikeCount(like.VideoId)
	videoLikeDao.Save(like)
	return nil
}

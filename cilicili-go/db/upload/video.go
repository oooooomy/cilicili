package upload

import (
	"cilicili-go/model/entity"
	"github.com/go-basic/uuid"
)

type BaseVideoDao interface {
	Save(video *entity.Video)
	FindById(id string) *entity.Video
	AddVideoCommentCount(id string)
	AddVideoLikeCount(id string)
	FindAll() []*entity.Video
	FindAllByMusicId(id string) []*entity.Video
	FindAllByUserId(id string) []*entity.Video
	FindAllByStatus(status bool) []*entity.Video
	SetStatusById(id string, status bool)
}

type VideoDao struct{}

func (v *VideoDao) Save(video *entity.Video) {
	video.Id = uuid.New()
	conn.Table("t_video").Save(video)
}

func (v *VideoDao) FindById(id string) *entity.Video {
	video := new(entity.Video)
	conn.Table("t_video").Find(video, "id = ?", id)
	return video
}

func (v *VideoDao) FindAll() (videos []*entity.Video) {
	conn.Table("t_video").Find(&videos)
	return
}

func (v *VideoDao) FindAllByMusicId(id string) (videos []*entity.Video) {
	conn.Table("t_video").Find(&videos, "music_id = ?", id)
	return
}

func (v *VideoDao) FindAllByUserId(id string) (videos []*entity.Video) {
	conn.Table("t_video").Find(&videos, "user_id = ?", id)
	return
}

func (v *VideoDao) FindAllByStatus(status bool) (videos []*entity.Video) {
	conn.Table("t_video").Find(&videos, "status = ?", status)
	return
}

func (v *VideoDao) SetStatusById(id string, status bool) {
	conn.Table("t_video").Where("id = ?", id).Update("status", status)
}

func (v *VideoDao) AddVideoCommentCount(id string) {
	conn.Exec("update t_video set comment_count = comment_count+1 where id = ?", id)
}

func (v *VideoDao) AddVideoLikeCount(id string) {
	conn.Exec("update t_video set like_count = like_count+1 where id = ?", id)
}

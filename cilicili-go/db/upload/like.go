package upload

import (
	"cilicili-go/model/entity"
	"github.com/go-basic/uuid"
)

type BaseVideoLikeDao interface {
	Save(like *entity.VideoLike)
	ExistsByUserIdAndVideoId(userid string, musicId string) bool
	FindAllByUserId(id string) []*entity.VideoLike
}

func (v *VideoLikeDao) Save(like *entity.VideoLike) {
	like.Id = uuid.New()
	conn.Table("t_video_like").Save(like)
}

func (v *VideoLikeDao) ExistsByUserIdAndVideoId(userid string, videoId string) bool {
	like := new(entity.VideoLike)
	conn.Table("t_video_like").Where("user_id = ? AND video_id = ?", userid, videoId).Find(like)
	return like.Id != ""
}

func (v *VideoLikeDao) FindAllByUserId(id string) (likes []*entity.VideoLike) {
	conn.Table("t_video_like").Find(&likes, "user_id = ?", id)
	return
}

type VideoLikeDao struct {
}

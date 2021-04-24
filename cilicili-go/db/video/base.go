package video

import (
	"cilicili-go/model/entity"
	"github.com/go-basic/uuid"
)

//Video数据库操作接口
type BaseVideoInterface interface {
	Save(video *entity.Video)
	Update(video *entity.Video)
	Delete(id string)
	FindAll() []*entity.Video
	FindByUserId(id string) []*entity.Video
	FindByMusicId(id string) []*entity.Video
}

//使用空结构体 实现接口
type BaseVideoDao struct{}

func (v *BaseVideoDao) Save(video *entity.Video) {
	video.Id = uuid.New()
	conn.Save(video)
}

func (v *BaseVideoDao) Update(video *entity.Video) {
	conn.Where(video, "id = ?", video.Id).Update(video)
}

func (v *BaseVideoDao) Delete(id string) {
	conn.Delete(&entity.Video{}, "id = ?", id)
}

func (v *BaseVideoDao) FindAll() (videos []*entity.Video) {
	conn.Find(&videos)
	for i := 0; i < len(videos); i++ {
		conn.Find(&videos[i].User, "id = ?", videos[i].UserId)
		conn.Find(&videos[i].Music, "id = ?", videos[i].MusicId)
	}
	return
}

func (v *BaseVideoDao) FindByUserId(id string) (videos []*entity.Video) {
	conn.Find(&videos, "user_id = ?", id)
	return
}

func (v *BaseVideoDao) FindByMusicId(id string) (videos []*entity.Video) {
	conn.Find(&videos, "music_id = ?", id)
	return
}

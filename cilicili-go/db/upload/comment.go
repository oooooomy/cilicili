package upload

import (
	"cilicili-go/model/entity"
	"github.com/go-basic/uuid"
	"time"
)

type BaseCommentDao interface {
	Save(comment *entity.Comment)
	FindAllByVideoId(id string) []*entity.Comment
}

type CommentDao struct{}

func (c *CommentDao) Save(comment *entity.Comment) {
	comment.Id = uuid.New()
	comment.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	conn.Table("t_comment").Save(comment)
}

func (c *CommentDao) FindAllByVideoId(id string) (comments []*entity.Comment) {
	conn.Table("t_comment").Find(&comments, "video_id = ?", id)
	return
}

package upload

import (
	db "cilicili-go/db/upload"
	"cilicili-go/model/entity"
)

type BaseCommentService interface {
	Create(comment *entity.Comment)
	FindAllByVideoId(id string) []*entity.Comment
}

type CommentService struct{}

var commentDao = new(db.CommentDao)

func (c CommentService) Create(comment *entity.Comment) {
	videoDao.AddVideoCommentCount(comment.VideoId)
	commentDao.Save(comment)
}

func (c CommentService) FindAllByVideoId(id string) []*entity.Comment {
	return commentDao.FindAllByVideoId(id)
}

package account

import (
	"cilicili-go/model/entity"
	"github.com/go-basic/uuid"
)

type BaseFollowDao interface {
	Save(follow *entity.Follow)
	Delete(from string, to string)
	FindAllByFromId(id string) []*entity.Follow
	ExistsByFromUserIdAndToUserId(from string, to string) bool
}

type FollowDao struct{}

func (f *FollowDao) Save(follow *entity.Follow) {
	follow.Id = uuid.New()
	conn.Table("t_follow").Save(follow)
}

func (f *FollowDao) Delete(from string, to string) {
	conn.Table("t_follow").Delete(entity.Follow{}, "from_user_id = ? AND to_user_id = ?", from, to)
}

func (f *FollowDao) FindAllByFromId(id string) (follows []*entity.Follow) {
	conn.Table("t_follow").Find(&follows, "from_user_id = ?", id)
	return
}

func (f *FollowDao) ExistsByFromUserIdAndToUserId(from string, to string) bool {
	follow := new(entity.Follow)
	conn.Table("t_follow").Where("from_user_id = ? AND to_user_id = ?", from, to).Find(follow)
	return follow.Id != ""
}

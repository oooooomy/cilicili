package account

import (
	"cilicili-go/db/account"
	"cilicili-go/model/entity"
	"errors"
)

type BaseFollowService interface {
	Create(follow *entity.Follow) error
	Delete(from string, to string)
	IsFollow(from string, to string) bool
	FindAllByFromId(id string) []*entity.Follow
}

var followDao = new(account.FollowDao)

func (f *FollowService) Create(follow *entity.Follow) error {
	if followDao.ExistsByFromUserIdAndToUserId(follow.FromUserId, follow.ToUserId) {
		return errors.New("已经添加了关注")
	}
	userDao.AddFollowCount(follow.FromUserId)
	userDao.AddFansCount(follow.ToUserId)
	followDao.Save(follow)
	return nil
}

func (f *FollowService) IsFollow(from string, to string) bool {
	return followDao.ExistsByFromUserIdAndToUserId(from, to)
}

func (f *FollowService) Delete(from string, to string) {
	userDao.DecFollowCount(from)
	userDao.DecFansCount(to)
	followDao.Delete(from, to)
}

func (f *FollowService) FindAllByFromId(id string) []*entity.Follow {
	return followDao.FindAllByFromId(id)
}

type FollowService struct{}

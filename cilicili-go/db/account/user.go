package account

import (
	"cilicili-go/model/entity"
)

type BaseUserDaoInterface interface {
	Save(user *entity.User)
	Update(user *entity.User)
	FindById(id string) *entity.User
	AddFollowCount(id string)
	AddFansCount(id string)
	DecFollowCount(id string)
	DecFansCount(id string)
}

type BaseUserDao struct{}

func (dao *BaseUserDao) Save(user *entity.User) {
	conn.Table("t_user").Save(user)
}

func (dao *BaseUserDao) Update(user *entity.User) {
	conn.Table("t_user").Updates(user)
}

func (dao *BaseUserDao) FindById(id string) *entity.User {
	user := &entity.User{}
	conn.Table("t_user").Find(user, "id = ?", id)
	return user
}

func (dao *BaseUserDao) AddFollowCount(id string) {
	conn.Exec("update t_user set follow_count = follow_count+1 where id = ?", id)
}

func (dao *BaseUserDao) AddFansCount(id string) {
	conn.Exec("update t_user set fans_count = fans_count+1 where id = ?", id)
}

func (dao *BaseUserDao) DecFollowCount(id string) {
	conn.Exec("update t_user set follow_count = follow_count-1 where id = ?", id)
}

func (dao *BaseUserDao) DecFansCount(id string) {
	conn.Exec("update t_user set fans_count = fans_count-1 where id = ?", id)
}

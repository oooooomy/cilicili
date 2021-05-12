package account

import (
	"cilicili-go/model/entity"
)

type BaseAdminDaoInterface interface {
	Save(admin *entity.Admin)
	Delete(id string)
	FindAll() []*entity.Admin
	FindByEmail(email string) *entity.Admin
}

type BaseAdminDao struct{}

func (dao *BaseAdminDao) Save(admin *entity.Admin) {
	conn.Table("t_admin").Save(admin)
}

func (dao *BaseAdminDao) Delete(id string) {
	conn.Table("t_admin").Delete(&entity.Admin{}, "id = ?", id)
}

func (dao *BaseAdminDao) FindAll() (admins []*entity.Admin) {
	conn.Table("t_admin").Find(&admins)
	return
}

func (dao *BaseAdminDao) FindByEmail(email string) *entity.Admin {
	admin := &entity.Admin{}
	conn.Table("t_admin").Where("email = ?", email).Find(admin)
	return admin
}


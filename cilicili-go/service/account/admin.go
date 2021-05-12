package account

import (
	"cilicili-go/client"
	"cilicili-go/db/account"
	"cilicili-go/model/entity"
	"errors"
)

type AdminServiceInterface interface {
	Create(admin *entity.Admin)
	FindAll() *[]entity.Admin
	Delete(id string)
	LoginByPassword(email string, password string) (*entity.Admin, error)
	LoginByEmailCode(email string, code string) (*entity.Admin, error)
}

type AdminService struct{}

var adminDao = new(account.BaseAdminDao)

func (a *AdminService) Create(admin *entity.Admin) {
	adminDao.Save(admin)
}

func (a *AdminService) FindAll() []*entity.Admin {
	return adminDao.FindAll()
}

func (a *AdminService) Delete(id string) {
	adminDao.Delete(id)
}

func (a *AdminService) LoginByPassword(email string, password string) (*entity.Admin, error) {
	admin := adminDao.FindByEmail(email)
	if admin.Id == "" {
		return nil, errors.New("邮箱或密码错误")
	}
	if admin.Password != password {
		return nil, errors.New("邮箱或密码错误")
	}
	return admin, nil
}

func (a *AdminService) LoginByEmailCode(email string, code string) (*entity.Admin, error) {
	admin := adminDao.FindByEmail(email)
	if admin.Id == "" {
		return nil, errors.New("身份验证失败")
	}
	emailClient := &client.EmailServiceClient{}
	err := emailClient.CheckCode(email, code)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

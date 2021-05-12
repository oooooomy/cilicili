package account

import (
	"cilicili-go/conf"
	"cilicili-go/db/account"
	"cilicili-go/model/dto"
	"cilicili-go/model/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BaseUserService interface {
	Login(code string) (*entity.User, error)
	Update(user *entity.User)
	GetInfo(id string) *entity.User
}

type UserService struct{}

var (
	userDao = new(account.BaseUserDao)
)

func (us *UserService) Login(code string) (*entity.User, error) {
	resp, err := http.Get("https://api.weixin.qq.com/sns/jscode2session?appid=" +
		conf.Config.App.AppId + "&secret=" + conf.Config.App.AppSecret +
		"&js_code=" + code + "&grant_type=authorization_code")
	if err != nil {
		fmt.Println("http request error: ", err.Error())
		return nil, err
	}
	loginDto := &dto.UserLoginDto{}
	data, _ := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(data, loginDto); err != nil {
		fmt.Println("json Unmarshal error: ", err.Error())
		return nil, err
	}
	user := userDao.FindById(loginDto.OpenId)
	if user.Id == "" {
		user.Id = loginDto.OpenId
		userDao.Save(user)
	}
	return user, nil
}

func (us *UserService) Update(user *entity.User) {
	userDao.Update(user)
}

func (us *UserService) GetInfo(id string) *entity.User {
	return userDao.FindById(id)
}

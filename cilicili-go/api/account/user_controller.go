package account

import (
	"cilicili-go/model/entity"
	"cilicili-go/model/support"
	"cilicili-go/service/account"
	"github.com/kataras/iris/v12"
)

var (
	userService  = new(account.UserService)
)

func userLogin(c iris.Context) {
	code := c.URLParam("code")
	if code == "" {
		_, _ = c.JSON(support.Error(400, "request param code missing"))
		return
	}
	user, err := userService.Login(code)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.SuccessWithData(user))
}

func userUpdate(c iris.Context) {
	user := &entity.User{}
	err := c.ReadJSON(user)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	userService.Update(user)
	_, _ = c.JSON(support.SuccessWithData(user))
}

func getUserInfo(c iris.Context) {
	id := c.Params().Get("id")
	user := userService.GetInfo(id)
	_, _ = c.JSON(support.SuccessWithData(user))
}


package account

import (
	"cilicili-go/model/entity"
	"cilicili-go/model/support"
	"cilicili-go/service/account"
	"github.com/kataras/iris/v12"
)

var (
	adminService = new(account.AdminService)
)

func adminLoginByPassword(c iris.Context) {
	admin := &entity.Admin{}
	err := c.ReadJSON(admin)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	if admin, err = adminService.LoginByPassword(admin.Email, admin.Password); err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.SuccessWithData(admin))
}

func adminLoginByEmail(c iris.Context) {
	email := c.URLParam("email")
	value := c.URLParam("value")
	admin, err := adminService.LoginByEmailCode(email, value)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.SuccessWithData(admin))
}

func findAllAdmin(c iris.Context) {
	_, _ = c.JSON(support.SuccessWithData(adminService.FindAll()))
}

func createAdmin(c iris.Context) {
	admin := &entity.Admin{}
	adminService.Create(admin)
	_, _ = c.JSON(support.Success())
}

func deleteAdmin(c iris.Context) {
	adminService.Delete(c.Params().Get("id"))
	_, _ = c.JSON(support.Success())
}

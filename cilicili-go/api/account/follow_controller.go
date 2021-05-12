package account

import (
	"cilicili-go/model/entity"
	"cilicili-go/model/support"
	"cilicili-go/service/account"
	"github.com/kataras/iris/v12"
)

var followService = new(account.FollowService)

func createFollow(c iris.Context) {
	f := new(entity.Follow)
	err := c.ReadJSON(f)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	err = followService.Create(f)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.Success())
}

func deleteFollow(c iris.Context) {
	followService.Delete(c.Params().Get("from"), c.Params().Get("to"))
	_, _ = c.JSON(support.Success())
}

func getUserFollowList(c iris.Context) {
	follows := followService.FindAllByFromId(c.Params().Get("id"))
	_, _ = c.JSON(support.SuccessWithData(follows))
}

func isFollow(c iris.Context) {
	ok := followService.IsFollow(c.Params().Get("from"), c.Params().Get("to"))
	_, _ = c.JSON(support.SuccessWithData(ok))
}

package room

import (
	"cilicili-go/conf"
	"cilicili-go/model/entity"
	"cilicili-go/model/support"
	"cilicili-go/service/room"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

var roomService = new(room.BaseRoomService)

func up(c iris.Context) {
	id := c.Params().Get("id")
	r := roomService.FindById(id)
	if r.Id == "" {
		_, _ = c.JSON(support.Error(400, "请求参数异常"))
		return
	}
	resp, err := http.Get("http://" + conf.Config.Live.Address + ":8090/control/get?room=" + id)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	type result struct {
		Status int    `json:"status"`
		Data   string `json:"data"`
	}
	apiResult := new(result)
	err = json.Unmarshal(data, apiResult)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	r.Token = apiResult.Data
	r.Url = "http://" + conf.Config.Live.Address + ":7002/cilicili/" + id + ".m3u8"
	roomService.Create(r)
	_, _ = c.JSON(support.SuccessWithData(apiResult.Data))
}

func down(c iris.Context) {

}

func findById(c iris.Context) {
	id := c.Params().Get("id")
	r := roomService.FindById(id)
	if r.Id == "" {
		_, _ = c.JSON(support.Error(400, "不存在的直播间"))
		return
	}
	_, _ = c.JSON(support.SuccessWithData(r))
}

func findByLiving(c iris.Context) {
	living := c.Params().Get("living")
	b, err := strconv.ParseBool(living)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.SuccessWithData(roomService.FindByLiving(b)))
}

func findByStatus(c iris.Context) {
	status := c.Params().Get("status")
	i, err := strconv.Atoi(status)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.SuccessWithData(roomService.FindByStatus(i)))
}

func saveRoom(c iris.Context) {
	r := new(entity.Room)
	err := c.ReadJSON(r)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	roomService.Create(r)
	_, _ = c.JSON(support.Success())
}

func findAllRoom(c iris.Context) {
	_, _ = c.JSON(support.SuccessWithData(roomService.FindAll()))
}

func wsHandler(c iris.Context) {
	userId := c.URLParam("userId")
	roomId := c.URLParam("roomId")
	logrus.Info("userId :", userId)
	logrus.Info("roomId :", roomId)
	ws := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	conn, err := ws.Upgrade(c.ResponseWriter(), c.Request(), nil)
	if err != nil {
		logrus.Error("websocket conn err: ", err.Error())
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	logrus.Info("conn :", conn)
}

package email

import (
	"cilicili-go/model/support"
	"cilicili-go/service/email"
	"github.com/kataras/iris/v12"
)

type Message struct {
	Content string
	To      string
}

var emailService = new(email.BaseEmailService)

func sendSimpleEmail(c iris.Context) {
	m := &Message{}
	err := c.ReadJSON(m)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	err = emailService.SendEmail(m.To, m.Content)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.Success())
}

func getEmailCode(c iris.Context) {
	to := c.URLParam("email")
	if to == "" {
		_, _ = c.JSON(support.Error(400, "sent to email missing"))
		return
	}
	err := emailService.SendCode(to)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.Success())
}

func checkEmailCode(c iris.Context) {
	e := c.URLParam("email")
	value := c.URLParam("value")
	err := emailService.CheckCode(e, value)
	if err != nil {
		_, _ = c.JSON(support.Error(400, err.Error()))
		return
	}
	_, _ = c.JSON(support.Success())
}

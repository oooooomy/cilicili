package email

import (
	"cilicili-go/conf"
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"sync"
	"time"
)

type BaseEmailServiceInterface interface {
	SendEmail(email string, content string) error
	SendCode(email string) error
	CheckCode(email string, value string) error
}

type BaseEmailService struct {
	codeMap sync.Map
}

func (b *BaseEmailService) SendEmail(email string, content string) error {
	return send(email, content)
}

func (b *BaseEmailService) SendCode(email string) error {
	//检查是否存在验证码
	code, ok := codeMap[email]
	if ok && code.exp > time.Now().Unix() {
		return errors.New("请勿频繁发送验证码")
	}
	value := randomCode()
	c := &emailCode{
		value: value,
		exp:   time.Now().Unix() + expTime,
	}
	codeMap[email] = c
	return send(email, "【cilicili】本次的验证码为："+value+"，请妥善保存验证码")
}

func (b *BaseEmailService) CheckCode(email string, value string) error {
	code, ok := codeMap[email]
	if !ok {
		return errors.New("不存在的邮箱")
	}
	if code.exp < time.Now().Unix() {
		return errors.New("验证码已过期")
	}
	if code.value != value {
		return errors.New("验证码错误，请仔细检查邮箱")
	}
	delete(codeMap, email)
	return nil
}

func send(to string, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", conf.Config.Mail.Username)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "cilicili")
	m.SetBody("text/html", content)
	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer(conf.Config.Mail.Host, conf.Config.Mail.Port,
		conf.Config.Mail.Username, conf.Config.Mail.Password)
	// 发送邮件
	return d.DialAndSend(m)
}

func randomCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

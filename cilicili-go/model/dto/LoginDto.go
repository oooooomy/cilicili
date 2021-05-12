package dto

type UserLoginDto struct {
	SessionKey string `json:"session_key"`
	OpenId     string `json:"openid"`
}

package client

import (
	"cilicili-go/model/support"
	"encoding/json"
	"errors"
	"net/http"
)

type EmailServiceClient struct{}

func (c *EmailServiceClient) CheckCode(email string, value string) error {
	url := "/check?email=" + email + "&value=" + value
	bytes, err := callServices("email-service", http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	result := &support.ResponseResult{}
	err = json.Unmarshal(bytes, result)
	if err != nil {
		return err
	}
	if result.Code != 200 {
		return errors.New(result.Msg)
	}
	return nil
}

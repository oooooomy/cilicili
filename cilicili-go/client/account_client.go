package client

import (
	"cilicili-go/model/entity"
	"encoding/json"
	"net/http"
)

type AccountServiceClient struct{}

func (a *AccountServiceClient) FindUserFollow(userId string) (follows []*entity.Follow) {
	url := "/follow/from/" + userId
	bytes, err := callServices("account-service", http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	type Result struct {
		Code int              `json:"code"`
		Data []*entity.Follow `json:"data"`
	}
	result := new(Result)
	err = json.Unmarshal(bytes, result)
	if err != nil {
		panic(err)
	}
	return result.Data
}

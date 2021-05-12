package client

import (
	"errors"
	eureka "github.com/xuanbo/eureka-client"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// callServices 服务调用
func callServices(serviceName string, method string, reqUrl string, body io.Reader) ([]byte, error) {
	var instances []eureka.Instance
	//获取服务实例
	for _, application := range eurekaClient.Applications.Applications {
		//eureka api返回的服务名为大写
		if strings.EqualFold(serviceName, application.Name) {
			instances = application.Instances
		}
	}
	if len(instances) == 0 {
		return nil, errors.New("The " + serviceName + " instance does not exist in eureka ")
	}
	//生成随机数 随机访问该服务下的所有实例
	serviceInstance := instances[doBalance(len(instances))]
	//此次请求的最终路径
	requestAddress := serviceInstance.HomePageURL + reqUrl
	//新建http request
	request, err := http.NewRequest(method, requestAddress, body)
	if err != nil {
		return nil, err
	}
	//新建http client
	httpClient := new(http.Client)
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func doBalance(len int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len)
}

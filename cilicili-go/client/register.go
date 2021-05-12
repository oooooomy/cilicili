package client

import (
	eureka "github.com/xuanbo/eureka-client"
)

var eurekaClient *eureka.Client

func RegisterToEureka(eurekaAddr string, port int, appName string) {
	// create eureka client
	eurekaClient = eureka.NewClient(&eureka.Config{
		DefaultZone:           eurekaAddr,
		App:                   appName,
		Port:                  port,
		RenewalIntervalInSecs: 10,
		DurationInSecs:        30,
	})
	// start client, register、heartbeat、refresh
	eurekaClient.Start()
}

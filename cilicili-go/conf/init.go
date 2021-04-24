package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	OSS    OssConfiguration
	DB     Database
	Eureka Eureka
}

var Config = Configuration{}

func init() {
	file, err := ioutil.ReadFile("../config.yaml")
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(file, &Config); err != nil {
		panic(err)
	}
}

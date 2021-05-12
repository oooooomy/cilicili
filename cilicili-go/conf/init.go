package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	OSS    OssConfiguration
	DB     Database
	Eureka Eureka
	App    App
	Mail   Mail
	Upload Upload
	Live   Live
}

type Database struct {
	Type string `yaml:"type"`
	Url  string `yaml:"url"`
}

type Eureka struct {
	Address string `yaml:"address"`
}

type App struct {
	AppId     string `yaml:"appId"`
	AppSecret string `yaml:"appSecret"`
}

type OssConfiguration struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
}

type Mail struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Live struct {
	Address string `yaml:"address"`
}

type Upload struct {
	Path   string `yaml:"path"`
	Ffmpeg string `yaml:"ffmpeg"`
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

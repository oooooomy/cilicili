package service

import (
	"cilicili-go/conf"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-basic/uuid"
	"io"
	"os"
)

var bucket *oss.Bucket

func LoadOssClient() {
	// 创建OSSClient实例。
	client, err := oss.New(conf.Config.OSS.Endpoint, conf.Config.OSS.AccessKeyId, conf.Config.OSS.AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 获取存储空间。
	bucket, err = client.Bucket(conf.Config.OSS.BucketName)
	if err != nil {
		panic(err)
	}
}

type UploadService interface {
	Upload(reader io.Reader, fileType string) (string, error)
}

type UploadServiceImpl struct{}

func (u *UploadServiceImpl) Upload(reader io.Reader, fileType string) (string, error) {
	// 指定存储类型为标准存储，缺省也为标准存储。
	storageType := oss.ObjectStorageClass(oss.StorageStandard)
	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)
	//保存的文件名设为uuid
	name := uuid.New() + "." + fileType
	err := bucket.PutObject(name, reader, storageType, objectAcl)
	if err != nil {
		return "", err
	}
	url := "https://" + conf.Config.OSS.BucketName + "." + conf.Config.OSS.Endpoint + "/" + name
	return url, nil
}

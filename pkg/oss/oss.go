package oss

import (
	"file-server/pkg/logging"
	"file-server/pkg/setting"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var ossCli *oss.Client

// Client : 创建oss client对象
func Client() *oss.Client {
	if ossCli != nil {
		return ossCli
	}
	ossCli, err := oss.New("http://"+ setting.OssSetting.OSSEndpoint,
		setting.OssSetting.OSSAccessKeyID, setting.OssSetting.OSSAccessKeySecret)
	if err != nil {
		logging.Error(err)
		return nil
	}
	return ossCli
}

// Bucket : 获取bucket存储空间
func Bucket() *oss.Bucket {
	cli := Client()
	if cli != nil {
		bucket, err := cli.Bucket(setting.OssSetting.OSSBucket)
		if err != nil {
			logging.Error(err)
			return nil
		}
		return bucket
	}
	return nil
}

// GetSignURL : 获取临时签名访问URL
func GetSignURL(objName string) string {
	signedURL, err := Bucket().SignURL(objName, oss.HTTPGet, 3600)
	if err != nil {
		logging.Error(err)
		return ""
	}
	return signedURL
}

//获取公共访问URL
func GetPublicURL(objName string) string{
	return  fmt.Sprintf("http://%s.%s/%s",setting.OssSetting.OSSBucket,setting.OssSetting.OSSEndpoint,objName)
}
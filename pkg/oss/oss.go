package oss

import (
	"file-server/pkg/logging"
	"file-server/pkg/setting"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var ossCli *oss.Client

// Client : 创建oss client对象
func Client() *oss.Client {
	if ossCli != nil {
		return ossCli
	}
	ossCli, err := oss.New(setting.OssSetting.OSSEndpoint,
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

// DownloadURL : 临时授权下载url
func DownloadURL(objName string) string {
	signedURL, err := Bucket().SignURL(objName, oss.HTTPGet, 3600)
	if err != nil {
		logging.Error(err)
		return ""
	}
	return signedURL
}
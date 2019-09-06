package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type FileInfo struct {
	ID         int64     `gorm:"primary_key" json:"id"`
	FileSha1    string    `json:"file_sha1"`
	FileName    string    `json:"file_name"`
	FileSize    int64     `json:"file_size"`
	FileAddress string    `json:"file_address"`
	CategoryNo  string    `json:"category_no"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	Status      int       `json:"status"`
}

func ExistFileInfo(fileSha1 string) (bool,error){
	var fileInfo FileInfo
	err := db.Select("file_sha1").Where("file_sha1 = ?", fileSha1).First(&fileInfo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	return true, nil
}

func AddFileInfo(fileSha1 string,fileName string,fileSize int64,fileAddress string,categoryNo string) error{
	fileInfo := FileInfo{FileSha1:fileSha1,FileName:fileName,FileSize:fileSize,FileAddress:fileAddress,CategoryNo:categoryNo}
	if err := db.Create(&fileInfo).Error;err !=nil{
		return err
	}
	return nil
}

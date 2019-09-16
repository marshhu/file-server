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
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	Status      int       `json:"status"`
}

func GetFileInfo(fileSha1 string) (*FileInfo,error){
	var fileInfo FileInfo
	err := db.Where("file_sha1 = ?", fileSha1).First(&fileInfo).Error
	if err != nil{
		return nil,err
	}
	return &fileInfo,nil
}

func ExistFileInfo(fileSha1 string) bool{
	var fileInfo FileInfo
	err := db.Select("file_sha1").Where("file_sha1 = ?", fileSha1).First(&fileInfo).Error
	if err != nil || err != gorm.ErrRecordNotFound{
      if len(fileInfo.FileSha1) > 0{
      	return  true
	  } else{
	  	return  false
	  }
	}
	if err == gorm.ErrRecordNotFound{
	  return  false
	}
	return true
}

func AddFileInfo(fileSha1 string,fileName string,fileSize int64,fileAddress string) error{
	fileInfo := FileInfo{FileSha1:fileSha1,FileName:fileName,FileSize:fileSize,FileAddress:fileAddress,CreateAt:time.Now(),UpdateAt:time.Now(),Status:1}
	if err := db.Create(&fileInfo).Error;err !=nil{
		return err
	}
	return nil
}

func DeleteFileInfo(fileSha1 string) error{
	err := db.Where("file_sha1 = ?", fileSha1).Delete(FileInfo{}).Error
	return err
}

func UpdateFileInfo(fileSha1 string,fileName string,fileSize int64,fileAddress string) error{
	fileInfo := FileInfo{FileSha1:fileSha1,FileName:fileName,FileSize:fileSize,FileAddress:fileAddress,UpdateAt:time.Now()}
	err := db.Model(&fileInfo).Updates(FileInfo{FileName:fileInfo.FileName,FileSize:fileInfo.FileSize,FileAddress:fileInfo.FileAddress,UpdateAt:time.Now()}).Error
	return err
}
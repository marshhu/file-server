package models

import (
	"time"
)

type UserFile struct {
	ID         int64     `gorm:"primary_key" json:"id"`
	UserNo     string    `json:"user_no"`
	FileSha1   string    `json:"file_sha1"`
	FileName   string    `json:"file_name"`
	UploadAt   time.Time `json:"upload_at"`
	LastUpdate time.Time `json:"last_update"`
	Status     int       `json:"status"`
}

func AddUserFile(userNo string,fileSha1 string,fileName string) error{
	userFile := UserFile{UserNo:userNo,FileSha1:fileSha1,FileName:fileName,UploadAt:time.Now()}
	if err := db.Create(&userFile).Error;err != nil{
		return err
	}
	return nil
}
package models

import (
	"time"
)

type UserFile struct {
	Id         int64     `gorm:"primary_key" json:"id"`
	UserNo     string    `json:"user_no"`
	FileSha1   string    `json:"file_sha1"`
	FileName   string    `json:"file_name"`
	UploadAt   time.Time `json:"upload_at"`
	LastUpdate time.Time `json:"last_update"`
	Status     int       `json:"status"`
}

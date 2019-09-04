package models

import (
	"time"
)

type FileInfo struct {
	Id          int64     `gorm:"primary_key" json:"id"`
	FileSha1    string    `json:"file_sha1"`
	FileName    string    `json:"file_name"`
	FileSize    int64     `json:"file_size"`
	FileAddress string    `json:"file_address"`
	CategoryNo  string    `json:"category_no"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
	Status      int       `json:"status"`
}

package models

import "time"

type FileInfo struct {
	ID       int       `gorm:"primary_key" json:"id"`
	FileSha1 string    `json:"file_sha1"`
	FileName string    `json:"file_name"`
	FileSize int64     `json:"file_size"`
	FileAddr string    `json:"file_addr"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	Status   int       `json:"status"`
	Ext1     int       `json:"ext1"`
	Ext2     string    `json:"ext2"`
}

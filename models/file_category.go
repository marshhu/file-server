package models

type FileCategory struct {
	ID         int64     `gorm:"primary_key" json:"id"`
	FileSha1    string    `json:"file_sha1"`
	CategoryNo    string    `json:"category_no"`
}

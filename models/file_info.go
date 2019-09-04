package models

import (
	"time"
)

type FileInfo struct {
	Id          int64     `xorm:"pk autoincr BIGINT(20)"`
	FileSha1    string    `xorm:"not null comment('文件hash') unique CHAR(40)"`
	FileName    string    `xorm:"not null comment('文件名') VARCHAR(256)"`
	FileSize    int64     `xorm:"not null comment('文件大小') BIGINT(20)"`
	FileAddress string    `xorm:"comment('文件存储位置') TEXT"`
	CategoryNo  string    `xorm:"comment('所属分类') VARCHAR(64)"`
	CreateAt    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('创建时间') DATETIME"`
	UpdateAt    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('更新时间') DATETIME"`
	Status      int       `xorm:"default 1 comment('状态(1、可用 2、禁用 3、已删除)') TINYINT(4)"`
}

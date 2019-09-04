package models

import (
	"time"
)

type UserFile struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	UserNo     string    `xorm:"not null comment('用户编号') VARCHAR(64)"`
	FileSha1   string    `xorm:"not null comment('文件hash') VARCHAR(255)"`
	FileName   string    `xorm:"comment('文件名') VARCHAR(255)"`
	UploadAt   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('上传时间') DATETIME"`
	LastUpdate time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('最后更新时间') DATETIME"`
	Status     int       `xorm:"default 1 comment('状态(1、可用 2、禁用 3、已删除)') INT(11)"`
}

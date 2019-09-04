package models

import (
	"time"
)

type User struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	UserNo     string    `xorm:"not null comment('用户编号') unique VARCHAR(64)"`
	Account    string    `xorm:"not null comment('登录账号') unique VARCHAR(64)"`
	Pwd        string    `xorm:"not null comment('登录密码') VARCHAR(256)"`
	UserName   string    `xorm:"comment('用户名') unique VARCHAR(64)"`
	Phone      string    `xorm:"comment('手机') VARCHAR(64)"`
	Email      string    `xorm:"comment('邮箱') VARCHAR(128)"`
	SignupAt   time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('注册时间') DATETIME"`
	LastActive time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('最后活动时间') DATETIME"`
	Profile    string    `xorm:"TEXT"`
	Status     int       `xorm:"default 1 comment('状态(1、可用 2、禁用 3、已删除)') TINYINT(4)"`
}

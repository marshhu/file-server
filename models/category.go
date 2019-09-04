package models

type Category struct {
	Id            int64  `xorm:"pk autoincr BIGINT(20)"`
	CategoryNo    string `xorm:"not null comment('分类编号') VARCHAR(64)"`
	CategoryTitle string `xorm:"not null comment('分类标题') TINYTEXT"`
	CategoryDesc  string `xorm:"comment('分类描述') TEXT"`
}

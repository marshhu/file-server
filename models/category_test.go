package models

import (
	"file-server/pkg/setting"
	"testing"
)

func clearTables() {
	db.Exec("truncate category")
	db.Exec("truncate file_info")
	db.Exec("truncate user_file")
	db.Exec("truncate user")
}

func TestMain(m *testing.M) {
	setting.Setup("../conf/app.ini")
	Setup()
	clearTables()
	m.Run()
	clearTables()
}

// func TesCategoryWorkFlow(t *testing.T) {
// 	t.Run("Add", testAddCategory)
// }

func TestAddCategory(t *testing.T) {
	err := AddCategory("Go语言高级教程", "Go语言即时通讯实战教程，升值加薪必备")
	if err != nil {
		t.Errorf("Error of AddCategory:%v", err)
	}
}

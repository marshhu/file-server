package models

import (
	"testing"
)

func TestAddUser(t *testing.T) {
 	err := AddUser("admin","1111","15300153009","65673781@qq.com","")
 	if err != nil{
		t.Errorf("Error of AddUser:%v", err)
	}
}

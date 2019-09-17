package models

import (
	"testing"
)

func TestAddUser(t *testing.T) {
 	err := AddUser("15300153009","123456","admin","15300153009","65673781@qq.com")
 	if err != nil{
		t.Errorf("Error of AddUser:%v", err)
	}
}

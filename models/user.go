package models

import (
	"file-server/pkg/util"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID        int64     `gorm:"primary_key" json:"id"`
	UserNo     string    `json:"user_no"`
	Account    string    `json:"account"`
	Pwd        string    `json:"pwd"`
	UserName   string    `json:"user_name"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	SignupAt   time.Time `json:"signup_at"`
	LastActive time.Time `json:"last_active"`
	Profile    string    `json:"profile"`
	Status     int       `json:"status"`
}

func AddUser(account string,pwd string,userName string,phone string,email string) error{
	uid,_ := uuid.NewV4()
	user := User{UserNo:uid.String(),Account:account,Pwd:util.EncodeMD5(pwd),UserName:userName,Phone:phone,Email:email,SignupAt:time.Now(),LastActive:time.Now(),Status:1}
	err := db.Create(&user).Error
	if err != nil{
		return err
	}
	return nil
}

func CheckUser(account,password string) (bool,error){
	user := User{}
	err := db.Select("id").Where(User{Account: account, Pwd: password}).First(&user).Error
	if err != nil {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}
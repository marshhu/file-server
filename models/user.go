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
	user := User{UserNo:uid.String(),Account:account,Pwd:util.EncodeMD5(pwd),Phone:phone,Email:email,SignupAt:time.Now()}
	if err := db.Create(&user).Error;err != nil{
		return err
	}
	return nil
}


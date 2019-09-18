package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID        int64     `gorm:"primary_key" json:"id"`
	UserNo     string    `json:"user_no"`
	UserName   string    `json:"user_name"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	SignupAt   time.Time `json:"signup_at"`
	LastActive time.Time `json:"last_active"`
	Profile    string    `json:"profile"`
	Status     int       `json:"status"`
}

func AddUser(username string,password string,phone string,email string,profile string) error{
	uid,_ := uuid.NewV4()
	user := User{UserNo:uid.String(),UserName:username,Password:password,Phone:phone,Email:email,SignupAt:time.Now(),LastActive:time.Now(),Profile:profile,Status:1}
	err := db.Create(&user).Error
	if err != nil{
		return err
	}
	return nil
}

func CheckUser(username string,password string) (bool,error){
	user := User{}
	err := db.Select("id").Where(User{UserName:username, Password: password}).First(&user).Error
	if err != nil {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}
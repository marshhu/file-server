package models

import (
	"time"
)

type User struct {
	Id         int64     `gorm:"primary_key" json:"id"`
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

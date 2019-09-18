package user_service

import "file-server/models"

type User struct {
	UserName string
	Password string
}

func (a *User) Login() (bool, error) {
	return models.CheckUser(a.UserName, a.Password)
}
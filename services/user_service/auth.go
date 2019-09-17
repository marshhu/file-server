package user_service

import "file-server/models"

type Auth struct {
	Account string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckUser(a.Account, a.Password)
}
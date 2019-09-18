package user_service

import "file-server/models"

type User struct {
	UserName string
	Password string
}

func (a *User) Login() (bool, error) {
	return models.CheckUser(a.UserName, a.Password)
}

type UserInfo struct{
	UserName string
	Avatar string
}
func (a *User) GetUserInfo() (*UserInfo,error){
 user,err := models.GetUserInfo(a.UserName)
 return &UserInfo{UserName:user.UserName,Avatar:user.Avatar},err
}
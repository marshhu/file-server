package api

import (
	"file-server/pkg/app"
	"file-server/pkg/e"
	"file-server/pkg/logging"
	"file-server/pkg/util"
	"file-server/services/user_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//登录
type User struct {
	UserName    string `form:"username"`
	Password string `form:"password"`
}

// @Summary Login
// @Produce  json
// @Param username query string
// @Param password query string
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/user/login [get]
func LoginHandler(c *gin.Context) {
	appG := app.Gin{C: c}
	a := User{}
	err := c.ShouldBind(&a)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err.Error)
		return
	}


	authService := user_service.User{UserName: a.UserName, Password: a.Password}
	isExist, err := authService.Login()
	if err != nil {
		logging.Error(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(a.UserName, a.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

//获取用户信息
func GetInfoHandler(c *gin.Context){
	fmt.Printf("this is GetInfoHandler")
}

//退出
func LogoutHandler(c *gin.Context){
	fmt.Printf("this is LogoutHandler")
}


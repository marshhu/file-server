package util

import (
	"file-server/pkg/setting"
	"github.com/gin-gonic/gin"
)

// GetPage get page parameters
func GetPage(c *gin.Context) int {
	result := 0
	page := StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}

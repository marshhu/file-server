package routers

import (
	"file-server/middleware/cors"
	"file-server/middleware/jwt"
	"file-server/routers/api"
	v1 "file-server/routers/api/v1"

	_ "file-server/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// 允许使用跨域请求  全局中间件
	router.Use(cors.Cors())
	//静态文件
	router.Static("/assets", "./static")
    router.Static("/tmp",".tmp")

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//渲染html页面
	router.LoadHTMLGlob("views/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	apiv1 := router.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		file := apiv1.Group("/file")
		{
			file.POST("/", v1.UploadHandler)
			file.GET("/:filesha1", v1.DownloadHandler)
			file.GET("/", v1.GetAllHandler)
			file.DELETE("/:filesha1", v1.DeleteHandler)
		}
		category := apiv1.Group("/category")
		{
			category.GET("/", v1.GetCategoryHandler)
			category.POST("/", v1.AddCategoryHandler)
			category.DELETE("/:category_no",v1.DeleteCategoryHandler)
		}
	}

	user := router.Group("/api/user")
	{
		user.POST("/login",api.LoginHandler)
		user.GET("/getInfo",api.GetInfoHandler)
		user.POST("/logout",api.LogoutHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

package routers

import (
	v1 "file-server/routers/api/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//静态文件
	router.Static("/assets", "./static")

	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	//渲染html页面
	router.LoadHTMLGlob("views/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	// r.GET("/auth", api.GetAuth)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := router.Group("/api/v1")
	{
		file := apiv1.Group("/file")
		{
			file.POST("/", v1.UploadHandler)
			file.GET("/:filesha1", v1.DownloadHandler)
			file.GET("/", v1.GetAllHandler)
			file.DELETE("/:filesha1", v1.DeleteHandler)
		}

	}
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

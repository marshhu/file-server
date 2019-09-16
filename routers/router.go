package routers

import (
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
	//apiv1 := router.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	//{
	//	file := apiv1.Group("/file")
	//	{
	//		file.POST("/", v1.UploadHandler)
	//		file.GET("/:filesha1", v1.DownloadHandler)
	//		file.GET("/", v1.GetAllHandler)
	//		file.DELETE("/:filesha1", v1.DeleteHandler)
	//	}
	//	category := apiv1.Group("/category")
	//	{
	//		category.GET("/", v1.GetCategoryHandler)
	//		category.POST("/", v1.AddCategoryHandler)
	//		category.DELETE("/:category_no",v1.DeleteCategoryHandler)
	//	}
	//}
	apiv1 := router.Group("/api/v1")
	{
		file := apiv1.Group("/file")
		{
			file.POST("/", v1.UploadHandler)
			file.GET("/:filesha1", v1.DownloadHandler)
			file.GET("/url/:filesha1", v1.GetFileUrlHandler)
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

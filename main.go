package main

import (
	"file-server/handler"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var serverUrl = "localhost:8080"

func main() {

	// Creates a router without any middleware by default
	router := gin.New()

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
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

	file := router.Group("/file")
	{
		file.POST("/upload", handler.UploadHandler)
		file.GET("/download/:filesha1", handler.DownloadHandler)
		file.GET("/list/", handler.GetAllHandler)
		file.DELETE("/delete/:filesha1", handler.DeleteHandler)
	}
	router.Run(serverUrl)
}

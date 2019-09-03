package v1

import (
	"file-server/pkg/setting"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//上传
func UploadHandler(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	log.Println(file.Size)
	// Upload the file to specific dst.
	fileAddr := setting.AppSetting.ImageSavePath + file.Filename
	log.Println(fileAddr)
	err := c.SaveUploadedFile(file, fileAddr)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

//下载
func DownloadHandler(c *gin.Context) {
	fmt.Printf("this is DownloadHandler")
}

//获取全部
func GetAllHandler(c *gin.Context) {
	fmt.Printf("this is GetAllHandler")
}

//删除
func DeleteHandler(c *gin.Context) {
	fmt.Printf("this is DeleteHandler")
}

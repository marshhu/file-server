package v1

import (
	"file-server/models"
	"file-server/pkg/app"
	"file-server/pkg/e"
	"file-server/pkg/file"
	"file-server/pkg/oss"
	"file-server/pkg/util"
	"file-server/services/file_info_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// @Summary 上传文件
// @Produce  json
// @Router /api/v1/file [post]
func UploadHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	f, _ := c.FormFile("file")

	// Upload the file to specific dst.
	//fileAddr := setting.AppSetting.ImageSavePath + file.Filename
	//log.Println(fileAddr)
	//err := c.SaveUploadedFile(file, fileAddr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	src, _ := f.Open()
	defer src.Close()
	fileMeta := file_info_service.FileMeta{
		FileShal:util.FileSha1(src),
		FileName:f.Filename,
		FileSize:f.Size,
		//FileAddress:fileAddr,
	}
    bucket := oss.Bucket()
	// 文件写入OSS存储
	objectKey := fileMeta.FileShal + file.GetExt(f.Filename)
	// 游标重新回到文件头部
	src.Seek(0,0)
	err := bucket.PutObject(objectKey, src)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOADFILE_FAIL, nil)
	}
	fileMeta.FileAddress = objectKey
    result,_ :=	fileMeta.AddFileInfo()
    if result{
		appG.Response(http.StatusOK, e.SUCCESS, fileMeta)
	} else{
		os.Remove(fileMeta.FileAddress)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FILEINFO_FAIL, nil)
	}

}

//下载
func DownloadHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	fileSha1 := c.Param("filesha1")
	fileInfo,_ := models.GetFileInfo(fileSha1)
	appG.Response(http.StatusOK, e.SUCCESS, fileInfo)
}

func GetFileUrlHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	fileSha1 := c.Param("filesha1")
	fileInfo,err := models.GetFileInfo(fileSha1)
	if err != nil{
		appG.Response(http.StatusNotFound, e.SUCCESS, nil)
	}
	signedURL := oss.DownloadURL(fileInfo.FileAddress)
	appG.Response(http.StatusOK, e.SUCCESS, signedURL)
}

//获取全部
func GetAllHandler(c *gin.Context) {
	fmt.Printf("this is GetAllHandler")
}

//删除
func DeleteHandler(c *gin.Context) {
	fmt.Printf("this is DeleteHandler")
}

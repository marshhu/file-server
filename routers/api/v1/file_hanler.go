package v1

import (
	"file-server/models"
	"file-server/pkg/app"
	"file-server/pkg/e"
	"file-server/pkg/oss"
	"file-server/pkg/setting"
	"file-server/pkg/util"
	"file-server/services/file_info_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// @Summary 单文件上传
// @Router /api/v1/file/uploadSingle [post]
func UploadSingleHandler(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	f, _ := c.FormFile("file")
	//Upload the file to specific dst.
	fileAddr := setting.AppSetting.TmpFilePath + f.Filename;
	log.Println(fileAddr)
	err := c.SaveUploadedFile(f, fileAddr)
	if err != nil {
		log.Fatal(err)
	}
	src, _ := f.Open()
	defer src.Close()
	fileMeta := file_info_service.FileMeta{
		FileShal:util.FileSha1(src),
		FileName:f.Filename,
		FileSize:f.Size,
		//FileAddress:fileAddr,
	}
    //bucket := oss.Bucket()
	//// 文件写入OSS存储
	//objectKey := fileMeta.FileShal + file.GetExt(f.Filename)
	//// 游标重新回到文件头部
	//src.Seek(0,0)
	//err := bucket.PutObject(objectKey, src)
	//if err != nil {
	//	appG.Response(http.StatusInternalServerError, e.ERROR_UPLOADFILE_FAIL, nil)
	//}
	//fileMeta.FileName = objectKey  //文件名
	//fileMeta.FileAddress = oss.GetPublicURL(objectKey)  //访问URL
    result,_ :=	fileMeta.AddFileInfo()
    if result{
		appG.Response(http.StatusOK, e.SUCCESS, fileMeta)
	} else{
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FILEINFO_FAIL, nil)
	}
}

func UploadMultiHandler(c *gin.Context){
	var (
		appG = app.Gin{C: c}
	)
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)
		f, _ := c.FormFile("file")
		//Upload the file to specific dst.
		fileAddr := setting.AppSetting.TmpFilePath + f.Filename
		log.Println(fileAddr)
		err := c.SaveUploadedFile(f, fileAddr)
		if err != nil {
			log.Fatal(err)
		}
		src, _ := f.Open()
		src.Close()
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
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
	signedURL := oss.GetSignURL(fileInfo.FileAddress)
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

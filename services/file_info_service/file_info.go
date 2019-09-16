package file_info_service

import (
	"file-server/models"
	"file-server/pkg/logging"
)

type FileMeta struct{
	FileShal string
	FileName string
	FileSize int64
	FileAddress string
}
func (f *FileMeta)AddFileInfo() (bool,error){
   isExist:= models.ExistFileInfo(f.FileShal)
   if !isExist{
	  err := models.AddFileInfo(f.FileShal, f.FileName, f.FileSize, f.FileAddress)
	  if err != nil {
	  	  logging.Error(err)
		   return false,nil
	  }
   } else{
		err :=	 models.UpdateFileInfo(f.FileShal, f.FileName, f.FileSize, f.FileAddress)
		if err != nil {
			logging.Error(err)
			return false,nil
		}
	}
   return  true,nil
}
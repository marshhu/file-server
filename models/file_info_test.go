package models

import (
	"file-server/pkg/util"
	"os"
	"testing"
)

//func TestAddFileInfo(t *testing.T) {
//	fileAddress :=`upload/images/avtaar.jpg`
//	f,_:= os.Open(fileAddress)
//	fileSha1 := util.FileSha1(f)
//	fileName := f.Name()
//	fileSize,_ := file.GetSize(f)
//
//	err := AddFileInfo(fileSha1,fileName,fileSize,fileAddress)
//	if err != nil {
//		t.Errorf("Error of AddFileInfo:%v", err)
//	}
//}

func TestExistFileInfo(t *testing.T) {
	fileAddress :=`upload/images/avtaar.jpg`
	f,_:= os.Open(fileAddress)
	fileSha1 := util.FileSha1(f)

	isExit := ExistFileInfo(fileSha1)
	if isExit{
	  t.Fail()
	}
}

//func TestDeleteFileInfo(t *testing.T) {
//	fileAddress :=`upload/images/avtaar.jpg`
//	f,_:= os.Open(fileAddress)
//	fileSha1 := util.FileSha1(f)
//
//    err := DeleteFileInfo(fileSha1)
//	if err != nil {
//		t.Errorf("Error of TestDeleteFileInfo:%v", err)
//	}
//
//}


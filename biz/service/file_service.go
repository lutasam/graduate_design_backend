package service

import (
	"github.com/gin-gonic/gin"
	"github.com/lutasam/doctors/biz/bo"
	"github.com/lutasam/doctors/biz/common"
	"github.com/lutasam/doctors/biz/utils"
	"strings"
	"sync"
)

type FileService struct{}

var (
	fileService     *FileService
	fileServiceOnce sync.Once
)

func GetFileService() *FileService {
	fileServiceOnce.Do(func() {
		fileService = &FileService{}
	})
	return fileService
}

func (ins *FileService) UploadImage(c *gin.Context, req *bo.UploadImageRequest) (*bo.UploadImageResponse, error) {
	isCorrect, err := utils.IsCorrectImg(req.FileHeader)
	if !isCorrect {
		return nil, err
	}
	newFilename := utils.Uint64ToString(utils.GenerateFileID()) + "." + strings.Split(req.FileHeader.Filename, ".")[1]
	url := utils.GetConfigString("file.imgs.url") + newFilename
	err = c.SaveUploadedFile(req.FileHeader, utils.GetConfigString("file.imgs.address")+newFilename)
	if err != nil {
		return nil, common.FILEUPLOADERROR
	}
	return &bo.UploadImageResponse{ImageURL: url}, nil
}

func (ins *FileService) UploadFile(c *gin.Context, req *bo.UploadFileRequest) (*bo.UploadFileResponse, error) {
	newFilename := utils.Uint64ToString(utils.GenerateFileID()) + "." + strings.Split(req.FileHeader.Filename, ".")[1]
	err := c.SaveUploadedFile(req.FileHeader, utils.GetConfigString("file.files.address")+newFilename)
	if err != nil {
		return nil, common.FILEUPLOADERROR
	}
	return &bo.UploadFileResponse{Filename: newFilename}, nil
}

func (ins *FileService) DownloadFile(c *gin.Context, req *bo.DownloadFileRequest) (*bo.DownloadFileResponse, error) {
	filepath := utils.GetConfigString("file.files.address") + req.Filename
	if !utils.IsFileExist(filepath) {
		return nil, common.FILENOTEXIST
	}
	c.File(filepath)
	return &bo.DownloadFileResponse{}, nil
}

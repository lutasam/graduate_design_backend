package utils

import (
	"github.com/lutasam/doctors/biz/common"
	"mime/multipart"
	"os"
	"strings"
)

func IsCorrectImg(header *multipart.FileHeader) (bool, error) {
	fileType := strings.Split(header.Filename, ".")[1]
	if fileType != "png" && fileType != "jpeg" && fileType != "jpg" {
		return false, common.IMGFORMATERROR
	}
	if header.Size > common.MAXIMGSPACE {
		return false, common.IMGTOOLARGEERROR
	}
	return true, nil
}

func IsFileExist(filepath string) bool {
	_, err := os.Lstat(filepath)
	return !os.IsNotExist(err)
}

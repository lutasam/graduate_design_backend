package bo

import "mime/multipart"

type UploadImageRequest struct {
	FileHeader *multipart.FileHeader `json:"image"`
}

type UploadImageResponse struct {
	ImageURL string `json:"image_url"`
}

type UploadFileRequest struct {
	FileHeader *multipart.FileHeader `json:"file"`
}

type UploadFileResponse struct {
	Filename string `json:"filename"`
}

type DownloadFileRequest struct {
	Filename string `json:"filename" binding:"required"`
}

type DownloadFileResponse struct{}

package backend

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/file/upload" method:"post" mine:"multipart/form-data" tags:"文件处理" summary:"本地上传文件"`
	File   *ghttp.UploadFile `json:"file" dc:"选择上传文件"`
}

type FileUploadRes struct {
	Id   uint   `json:"id" dc:"文件id"`
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url" dc:"访问URL，可能只是URI"`
}

type FileUploadOssReq struct {
	g.Meta `path:"/oss/upload" method:"post" mine:"multipart/form-data" tags:"文件处理" summary:"oss上传文件"`
	File   *ghttp.UploadFile `json:"file" dc:"选择上传文件"`
}

type FileUploadOssRes struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"url" dc:"访问URL，可能只是URI"`
}

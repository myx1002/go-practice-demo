package controller

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"goframe-shop/api/backend"
	"goframe-shop/internal/model"
	"goframe-shop/internal/service"
	"goframe-shop/utility"
)

var File cFile

type cFile struct {
}

func (a *cFile) Upload(ctx context.Context, req *backend.FileUploadReq) (res *backend.FileUploadRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}

	result, err := service.File().FileUpload(ctx, model.FileUploadInput{
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}

	res = &backend.FileUploadRes{
		Id:   result.Id,
		Name: result.Name,
		Url:  result.Url,
	}
	return
}

func (a *cFile) UploadOssFile(ctx context.Context, req *backend.FileUploadOssReq) (res *backend.FileUploadOssRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择需要上传的文件")
	}

	name, url, err := utility.UploadOssFile(ctx, req.File)
	if err != nil {
		return nil, err
	}

	res = &backend.FileUploadOssRes{
		Name: name,
		Url:  url,
	}
	return
}

package file

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
)

type sFile struct {
}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

/*
*
1.定义图片上传位置
2.校验：上传位置是否正确、安全性校验：1分钟内只能10次
3.定义年月日 Ymd
4.入库
5.返回数据
*/
func (s *sFile) FileUpload(ctx context.Context, in model.FileUploadInput) (out *model.FileUploadOutput, err error) {
	// 1.定义图片上传位置
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		return nil, gerror.New("获取文件上传路劲失败！")
	}

	// 2.校验：上传位置是否正确、安全性校验：1分钟内只能10次
	count, err := dao.FileInfo.Ctx(ctx).
		Where(dao.FileInfo.Columns().UserId, ctx.Value(consts.CtxAdminId)).
		WhereGTE(dao.FileInfo.Columns().CreatedAt, gtime.Now().Add(-time.Minute)).
		Count()
	if err != nil {
		return nil, err
	}
	if count >= consts.FileMaxUploadCountMinute {
		return nil, gerror.New("请勿频繁上传文件！")
	}

	// 3.定义年月日 Ymd
	dateDirName := gtime.Now().Format("Ymd")
	fileName, err := in.File.Save(gfile.Join(uploadPath, dateDirName), in.RandomName)
	if err != nil {
		return nil, err
	}

	// 4.入库
	data := entity.FileInfo{
		UserId:       gconv.Int(ctx.Value(consts.CtxAdminId)),
		Name:         fileName,
		RealFileName: in.File.Filename,
		Src:          gfile.Join(uploadPath, dateDirName, fileName),
		Url:          "/upload/" + dateDirName + "/" + fileName,
	}
	id, err := dao.FileInfo.Ctx(ctx).Data(data).OmitEmpty().InsertAndGetId()
	if err != nil {
		return nil, err
	}

	// 5.返回数据
	return &model.FileUploadOutput{
		Id:   uint(id),
		Name: data.Name,
		Src:  data.Src,
		Url:  data.Url,
	}, nil
}

package login

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
	"goframe-shop/utility"
)

type sLogin struct{}

func init() {
	// 通过配置goland File Watcher自动生成service
	// https://goframe.org/pages/viewpage.action?pageId=49770772&preview=/49770772/49770777/watchers.xml
	service.RegisterLogin(New())
}

func New() *sLogin {
	return &sLogin{}
}

// 执行登录
func (s *sLogin) Login(ctx context.Context, in model.UserLoginInput) error {
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return gerror.New("该账号不存在！")
	}

	// 用于打印调试
	//gutil.Dump(utility.EncrypPassword(in.Password, adminInfo.UserSalt))

	if utility.EncrypPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return gerror.New("密码有误，请重新输入！")
	}

	if err := service.Session().SetUser(ctx, &adminInfo); err != nil {
		return err
	}
	// 自动更新上线
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		Id:      uint(adminInfo.Id),
		Name:    adminInfo.Name,
		IsAdmin: int8(adminInfo.IsAdmin),
	})
	return nil
}

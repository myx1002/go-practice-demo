package admin

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/gogf/gf/v2/util/gutil"

	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
	"goframe-shop/utility"

	"github.com/gogf/gf/v2/encoding/ghtml"

	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
)

type sAdmin struct{}

func init() {
	// 通过配置goland File Watcher自动生成service
	// https://goframe.org/pages/viewpage.action?pageId=49770772&preview=/49770772/49770777/watchers.xml
	service.RegisterAdmin(New())
}

func New() *sAdmin {
	return &sAdmin{}
}

// Create 创建内容
func (s *sAdmin) Create(ctx context.Context, in model.AdminCreateInput) (out model.AdminCreateOutput, err error) {
	// 不允许HTML代码
	if err = ghtml.SpecialCharsMapOrStruct(in); err != nil {
		return out, err
	}
	// 获取盐值 && 密码加密
	UserSalt := grand.S(10)
	in.Password = utility.EncrypPassword(in.Password, UserSalt)
	in.UserSalt = UserSalt

	lastInsertID, err := dao.AdminInfo.Ctx(ctx).Data(in).InsertAndGetId()
	if err != nil {
		return out, err
	}
	return model.AdminCreateOutput{AdminId: uint(lastInsertID)}, err
}

// Delete 删除
func (s *sAdmin) Delete(ctx context.Context, id uint) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除内容-软删除
		// 如果要硬删除，可加.Unscoped()来忽略时间特性
		_, err := dao.AdminInfo.Ctx(ctx).Where(g.Map{
			dao.AdminInfo.Columns().Id: id,
		}).Delete()
		return err
	})
}

// Update 修改
func (s *sAdmin) Update(ctx context.Context, in model.AdminUpdateInput) error {
	return dao.AdminInfo.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 不允许HTML代码
		if err := ghtml.SpecialCharsMapOrStruct(in); err != nil {
			return err
		}

		// 密码不为空时，修改密码
		if in.Password != "" {
			UserSalt := grand.S(10)
			in.Password = utility.EncrypPassword(in.Password, UserSalt)
			in.UserSalt = UserSalt
		}

		_, err := dao.AdminInfo.
			Ctx(ctx).
			Data(in).
			FieldsEx(dao.AdminInfo.Columns().Id). // 用于过滤Id字段，即Exclude
			Where(dao.AdminInfo.Columns().Id, in.Id).
			Update()
		return err
	})
}

// GetList 查询内容列表
func (s *sAdmin) GetList(ctx context.Context, in model.AdminGetListInput) (out *model.AdminGetListOutput, err error) {
	var (
		m = dao.AdminInfo.Ctx(ctx)
	)
	out = &model.AdminGetListOutput{
		Page: in.Page,
		Size: in.Size,
	}

	// 分配查询
	listModel := m.Page(in.Page, in.Size)
	// 排序方式
	listModel = listModel.OrderDesc(dao.AdminInfo.Columns().Id)

	// 执行查询
	var list []*entity.AdminInfo
	if err := listModel.Scan(&list); err != nil {
		return out, err
	}
	// 没有数据
	if len(list) == 0 {
		return out, nil
	}
	out.Total, err = m.Count()
	if err != nil {
		return out, err
	}
	// Admin
	if err := listModel.Scan(&out.List); err != nil {
		return out, err
	}
	return
}

func (s *sAdmin) GetUserByUserNamePassword(ctx context.Context, in model.UserLoginInput) map[string]interface{} {
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where("name", in.Name).Scan(&adminInfo)
	if err != nil {
		return nil
	}

	// 用于打印调试
	gutil.Dump("origin", utility.EncrypPassword(in.Password, adminInfo.UserSalt))
	gutil.Dump("pass", adminInfo.Password)
	if utility.EncrypPassword(in.Password, adminInfo.UserSalt) != adminInfo.Password {
		return nil
	} else {
		return g.Map{
			"id":   adminInfo.Id,
			"name": adminInfo.Name,
		}
	}
}

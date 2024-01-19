package cmd

import (
	"context"
	"strconv"
	"time"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"

	"goframe-shop/api/backend"
	"goframe-shop/api/frontend"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model/entity"
	"goframe-shop/utility"
	"goframe-shop/utility/response"
)

// 获取管理后台token
func getBackendAdminGToken() (gfAdminToken *gtoken.GfToken, err error) {
	gfAdminToken = &gtoken.GfToken{
		ServerName: "shop",
		CacheMode:  2, // gredis 需要配置redis配置
		//Timeout:         10 * 1000,
		LoginPath:        "/login",
		LoginBeforeFunc:  loginFunc,
		LoginAfterFunc:   loginAfterFunc,
		LogoutPath:       "/logout",
		AuthPaths:        g.SliceStr{}, // 拦截路劲
		AuthExcludePaths: g.SliceStr{}, // 不拦截路径
		MultiLogin:       true,         // 是否支持多端登录
		AuthAfterFunc:    authAfterFunc,
	}

	return
}

// 获取前台gtoken
func getFrontendGToken() (gfAdminToken *gtoken.GfToken, err error) {
	gfAdminToken = &gtoken.GfToken{
		ServerName: "shop",
		CacheMode:  2, // gredis 需要配置redis配置
		//Timeout:         10 * 1000,
		LoginPath:        "/login",
		LoginBeforeFunc:  loginFrontendFunc,
		LoginAfterFunc:   loginAfterFrontendFunc,
		LogoutPath:       "/logout",
		AuthPaths:        g.SliceStr{}, // 拦截路劲
		AuthExcludePaths: g.SliceStr{}, // 不拦截路径
		MultiLogin:       false,        // 是否支持多端登录
		AuthAfterFunc:    authAfterFrontendFunc,
	}

	return
}

func loginFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}
	ctx := context.TODO()
	adminInfo := entity.AdminInfo{}
	err := dao.AdminInfo.Ctx(ctx).Where(dao.AdminInfo.Columns().Name, name).Scan(&adminInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}

	if utility.EncrypPassword(password, adminInfo.UserSalt) != adminInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}

	// 返回值1:唯一标识
	// 返回值2:扩展参数user data,可在authAfterFunc的respData gtoken.Resp中获取
	return consts.GTokenAdminPrefix + strconv.Itoa(adminInfo.Id), adminInfo
}

func loginFrontendFunc(r *ghttp.Request) (string, interface{}) {
	name := r.Get("name").String()
	password := r.Get("password").String()

	if name == "" || password == "" {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}
	ctx := context.TODO()
	userInfo := entity.UserInfo{}
	err := dao.UserInfo.Ctx(ctx).Where(dao.UserInfo.Columns().Name, name).Scan(&userInfo)
	if err != nil {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}

	if utility.EncrypPassword(password, userInfo.UserSalt) != userInfo.Password {
		r.Response.WriteJson(gtoken.Fail(consts.ErrLoginFailMsg))
		r.ExitAll()
	}

	// 返回值1:唯一标识
	// 返回值2:扩展参数user data,可在authAfterFunc的respData gtoken.Resp中获取
	return consts.GTokenUserPrefix + strconv.Itoa(userInfo.Id), userInfo
}

// 自定义登录之后才会到这里
func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	// {
	//    Code: 0,
	//    Msg:  "success",
	//    Data: {
	//        "token":   "JxqZzeumBnmV8tsn0ybC+DVXirouyScYOQiNSEIux688av2hHewX4A5cjM4qonZ6",
	//        "userKey": "Admin:13",
	//        "uuid":    "d1e63353652f266b44b52a6ebc4113eb",
	//    },
	//}
	//g.Dump("respDataloginAfterFunc:", respData)
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {
		respData.Code = 1
		// 获取登录用户的id
		userKey := respData.Get("userKey").String()
		adminId := gstr.StrEx(userKey, consts.GTokenAdminPrefix)

		// 根据id获取登录用户信息
		adminInfo := &entity.AdminInfo{}
		err := dao.AdminInfo.Ctx(context.TODO()).WherePri(adminId).Scan(&adminInfo)
		if err != nil {
			return
		}

		// 查询用户角色对应的权限id
		var rolePermissionInfos []entity.RolePermissionInfo
		err = dao.RolePermissionInfo.Ctx(context.TODO()).WhereIn(dao.RolePermissionInfo.Columns().RoleId, g.Slice{adminInfo.RoleIds}).Scan(&rolePermissionInfos)
		if err != nil {
			return
		}
		permissionIds := g.Slice{}
		for _, info := range rolePermissionInfos {
			permissionIds = append(permissionIds, info.PermissionId)
		}

		// 查询权限
		var permissionInfos []entity.PermissionInfo
		err = dao.PermissionInfo.Ctx(context.TODO()).WhereIn(dao.PermissionInfo.Columns().Id, permissionIds).Scan(&permissionInfos)
		if err != nil {
			return
		}
		data := &backend.LoginRes{
			Type:        "Bearer",
			Token:       respData.GetString("token"),
			ExpireIn:    int(time.Hour * 24),
			IsAdmin:     adminInfo.IsAdmin,
			RoleIds:     adminInfo.RoleIds,
			Permissions: permissionInfos,
		}
		response.JsonExit(r, 0, "", data)
	}
}

// 自定义登录之后才会到这里
func loginAfterFrontendFunc(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		respData.Code = 0
		r.Response.WriteJson(respData)
		return
	} else {

		respData.Code = 1
		// 获取登录用户的id
		userKey := respData.Get("userKey").String()
		userId := gstr.StrEx(userKey, consts.GTokenUserPrefix)

		// 根据id获取登录用户信息
		userInfo := &entity.UserInfo{}
		err := dao.UserInfo.Ctx(context.TODO()).WherePri(userId).Scan(&userInfo)
		if err != nil {
			return
		}

		data := &frontend.LoginRes{
			Type:     "Bearer",
			Token:    respData.GetString("token"),
			ExpireIn: int(time.Hour * 24),
			UserId:   uint(userInfo.Id),
			Name:     userInfo.Name,
		}
		response.JsonExit(r, 0, "", data)
	}
}

// 认证后执行-记录用户信息
func authAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	//g.Dump("respDataauthAfterFunc:", respData)
	// {"Code":0,"Msg":"success","Data":{"createTime":1700488052777,"refreshTime":1700920052777,"userKey":"Admin:13","uuid":"b933618559124f25e4f6b8cc3a06bb34","data":{"Id":13,"Name":"吕霞","Password":"94f6ad86ed8c34e986220f7165d34a5b","RoleIds":"0","CreatedAt":"2023-11-16 08:11:30","UpdatedAt":"2023-11-17 08:16:46","DeletedAt":"","UserSalt":"pbpcxN91lZ","IsAdmin":1}}}
	var adminInfo entity.AdminInfo
	err := gconv.Struct(respData.GetString("data"), &adminInfo)
	if err != nil {
		response.Auth(r)
		return
	}

	if adminInfo.DeletedAt != nil {
		response.AuthBlack(r)
		return
	}

	// 把token里面值设置到上下文信息中
	r.SetCtxVar(consts.CtxAdminId, adminInfo.Id)
	r.SetCtxVar(consts.CtxAdminName, adminInfo.Name)
	r.SetCtxVar(consts.CtxAdminIsAdmin, adminInfo.IsAdmin)
	r.SetCtxVar(consts.CtxAdminRoleIds, adminInfo.RoleIds)

	// 继续执行
	r.Middleware.Next()
}

// 认证后执行-记录用户信息
func authAfterFrontendFunc(r *ghttp.Request, respData gtoken.Resp) {
	var userInfo entity.UserInfo
	err := gconv.Struct(respData.GetString("data"), &userInfo)
	if err != nil {
		response.Auth(r)
		return
	}

	if userInfo.DeletedAt != nil || userInfo.Status == consts.CtxUserStatusBlock {
		response.AuthBlack(r)
		return
	}

	// 把token里面值设置到上下文信息中
	r.SetCtxVar(consts.CtxUserId, userInfo.Id)
	r.SetCtxVar(consts.CtxUserName, userInfo.Name)
	r.SetCtxVar(consts.CtxUserAvatar, userInfo.Avatar)
	r.SetCtxVar(consts.CtxUserStatus, userInfo.Status)
	r.SetCtxVar(consts.CtxUserSign, userInfo.Sign)

	// 继续执行
	r.Middleware.Next()
}

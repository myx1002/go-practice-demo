package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"go-zero-demo/userapi/common/result"
	"go-zero-demo/userapi/internal/logic/user"
	"go-zero-demo/userapi/internal/svc"
	"go-zero-demo/userapi/internal/types"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.UserInfo(&req)
		result.HttpResult(r, w, resp, err)
	}
}

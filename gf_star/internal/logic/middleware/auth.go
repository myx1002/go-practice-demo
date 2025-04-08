package middleware

import (
	"gf_star/internal/consts"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func Auth(r *ghttp.Request) {
	// 从头部的Authorization获取token
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.JwtKey), nil
	})

	if err != nil || !token.Valid {
		r.Response.WriteStatus(http.StatusForbidden)
		r.Exit()
	}

	// 继续执行后续的请求处理
	r.Middleware.Next()
}

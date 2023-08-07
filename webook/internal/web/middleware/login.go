package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddleWareBuilder struct {
	pathSlice []string
}

func NewLoginMiddleWareBuilder() *LoginMiddleWareBuilder {
	return &LoginMiddleWareBuilder{}
}

func (l *LoginMiddleWareBuilder) IgnorePath(p string) *LoginMiddleWareBuilder {
	l.pathSlice = append(l.pathSlice, p)
	return l
}

func (l *LoginMiddleWareBuilder) Builder() gin.HandlerFunc {
	// todo 校验是否登录
	return func(ctx *gin.Context) {
		for _, v := range l.pathSlice {
			if v == ctx.Request.RequestURI {
				return
			}
		}

		sess := sessions.Default(ctx)
		if sess.Get("userId") == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

	}
}

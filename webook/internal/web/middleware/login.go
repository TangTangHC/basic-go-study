package middleware

import (
	"fmt"
	"github.com/TangTangHC/basic-go-study/webook/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
	"time"
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

		//sess := sessions.Default(ctx)
		//if sess.Get("userId") == nil {
		//	ctx.AbortWithStatus(http.StatusUnauthorized)
		//}

		jwtToken := ctx.GetHeader("Authorization")
		split := strings.Split(jwtToken, " ")
		if len(split) != 2 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		jwtToken = split[1]
		userClaims := &web.UserClaims{}
		token, err := jwt.ParseWithClaims(jwtToken, userClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte("TyrmfzW2KnkH0HRfIH6lzd5XsQtrM31O"), nil
		})
		if err != nil {
			fmt.Println(err)
		}
		if token == nil || !token.Valid || userClaims.UId == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if ctx.Request.UserAgent() != userClaims.UserAgent {
			// todo 监控这里
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// 是否刷新缓存
		if userClaims.ExpiresAt.Sub(time.Now()) < time.Minute*2 {
			userClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute))
			signedString, err := token.SignedString([]byte("TyrmfzW2KnkH0HRfIH6lzd5XsQtrM31O"))
			if err != nil {
				log.Println("刷新token失败")
			}
			ctx.Header("x-jwt-token", signedString)

		}
		ctx.Set("token", userClaims)
	}
}

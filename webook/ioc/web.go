package ioc

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/TangTangHC/basic-go-study/webook/internal/web"
	"github.com/TangTangHC/basic-go-study/webook/internal/web/middleware"
)

func InitWebServer(mdls []gin.HandlerFunc, handler *web.UserHandler) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)
	handler.RegisterHandler(server)
	return server
}

func InitMiddleWare() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		corsConfig(),
		middleware.NewLoginMiddleWareBuilder().
			IgnorePath("/users/signup").
			IgnorePath("/users/login").
			IgnorePath("/users/login_sms/code/send").
			IgnorePath("/users/login_sms").
			Builder(),
		sessionConfig(),
		// todo 限流
	}
}

func corsConfig() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"x-jwt-token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			//return strings.HasPrefix(origin, "http://localhost")
			return true
		},
	})
}

func sessionConfig() gin.HandlerFunc {
	//store := cookie.NewStore([]byte("secret"))
	store := memstore.NewStore([]byte("OxBC4Y5fWGcsGoTlQpYIr4HzDAcZTyk4kbUFK2MqSPyWKJJ9A3JrStYXOYktSb3B"),
		[]byte("PiZuMqQnEdbgsfHcYwBoduDtGbnaK7dj"))
	//store := memstore.NewStore([]byte("secret"))
	//store, _ := redis.NewStore(6, "tcp", config.Config.Redis.Addr, "", []byte("OxBC4Y5fWGcsGoTlQpYIr4HzDAcZTyk4kbUFK2MqSPyWKJJ9A3JrStYXOYktSb3B"),
	//	[]byte("PiZuMqQnEdbgsfHcYwBoduDtGbnaK7dj"))
	return sessions.Sessions("mysession", store)
}

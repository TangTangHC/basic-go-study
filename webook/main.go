package main

import (
	"github.com/TangTangHC/basic-go-study/webook/config"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/dao"
	"github.com/TangTangHC/basic-go-study/webook/internal/service"
	"github.com/TangTangHC/basic-go-study/webook/internal/web"
	"github.com/TangTangHC/basic-go-study/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func main() {
	db := initDB()

	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"x-jwt-token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			//return strings.HasPrefix(origin, "http://localhost")
			return true
		},
	}))
	//store := cookie.NewStore([]byte("secret"))
	//store := memstore.NewStore([]byte("OxBC4Y5fWGcsGoTlQpYIr4HzDAcZTyk4kbUFK2MqSPyWKJJ9A3JrStYXOYktSb3B"),
	//	[]byte("PiZuMqQnEdbgsfHcYwBoduDtGbnaK7dj"))
	//store := memstore.NewStore([]byte("secret"))
	store, _ := redis.NewStore(6, "tcp", config.Config.Redis.Addr, "", []byte("OxBC4Y5fWGcsGoTlQpYIr4HzDAcZTyk4kbUFK2MqSPyWKJJ9A3JrStYXOYktSb3B"),
		[]byte("PiZuMqQnEdbgsfHcYwBoduDtGbnaK7dj"))
	server.Use(sessions.Sessions("mysession", store))

	loginMiddleWareBuilder := middleware.NewLoginMiddleWareBuilder()
	server.Use(loginMiddleWareBuilder.IgnorePath("/users/signup").IgnorePath("/users/login").Builder())

	userHandler := initUserHandler(db)
	userHandler.RegisterHandler(server)
	server.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, " pong")
	})
	err := server.Run(":8080")
	if err != nil {
		return
	}
}

func initUserHandler(db *gorm.DB) *web.UserHandler {
	userDao := dao.NewUserDao(db)
	uRep := repository.NewUserRepository(userDao)
	srv := service.NewUserService(uRep)
	return web.NewUserHandler(srv)
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	db = db.Debug()
	if err != nil {
		panic("系统初始化错误")
	}
	return db
}

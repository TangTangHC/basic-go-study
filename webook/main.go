package main

import (
	"github.com/TangTangHC/basic-go-study/webook/internal/repository"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/dao"
	"github.com/TangTangHC/basic-go-study/webook/internal/service"
	"github.com/TangTangHC/basic-go-study/webook/internal/web"
	"github.com/TangTangHC/basic-go-study/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
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
			return strings.HasPrefix(origin, "http://localhost")
		},
	}))
	//store := cookie.NewStore([]byte("secret"))
	store := memstore.NewStore([]byte("OxBC4Y5fWGcsGoTlQpYIr4HzDAcZTyk4kbUFK2MqSPyWKJJ9A3JrStYXOYktSb3B"),
		[]byte("PiZuMqQnEdbgsfHcYwBoduDtGbnaK7dj"))
	//store := memstore.NewStore([]byte("secret"))
	//store, _ := redis.NewStore(6, "tcp", "localhost:6379", "", []byte("OxBC4Y5fWGcsGoTlQpYIr4HzDAcZTyk4kbUFK2MqSPyWKJJ9A3JrStYXOYktSb3B"),
	//	[]byte("PiZuMqQnEdbgsfHcYwBoduDtGbnaK7dj"))
	server.Use(sessions.Sessions("mysession", store))

	loginMiddleWareBuilder := middleware.NewLoginMiddleWareBuilder()
	server.Use(loginMiddleWareBuilder.IgnorePath("/users/signup").IgnorePath("/users/login").Builder())

	userHandler := initUserHandler(db)
	userHandler.RegisterHandler(server)
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
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/g_webook"))
	db = db.Debug()
	if err != nil {
		panic("系统初始化错误")
	}
	return db
}

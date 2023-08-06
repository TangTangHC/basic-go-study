package main

import (
	"github.com/TangTangHC/basic-go-study/webook/internal/repository"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/dao"
	"github.com/TangTangHC/basic-go-study/webook/internal/service"
	"github.com/TangTangHC/basic-go-study/webook/internal/web"
	"github.com/gin-contrib/cors"
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
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			return strings.HasPrefix(origin, "http://localhost")
		},
	}))
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
	if err != nil {
		panic("系统初始化错误")
	}
	return db
}

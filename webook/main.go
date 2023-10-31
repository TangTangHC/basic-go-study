package main

import (
	"github.com/TangTangHC/basic-go-study/webook/config"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository"
	redis2 "github.com/TangTangHC/basic-go-study/webook/internal/repository/cache/redis"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/dao"
	"github.com/TangTangHC/basic-go-study/webook/internal/service"
	"github.com/TangTangHC/basic-go-study/webook/internal/service/sms/memory"
	"github.com/TangTangHC/basic-go-study/webook/internal/web"
	"github.com/TangTangHC/basic-go-study/webook/internal/web/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func main() {
	db := initDB()
	rdb := initRedis()

	server := initWebServer()

	u := initUser(db, rdb)
	u.RegisterHandler(server)

	err := server.Run(":8080")
	if err != nil {
		return
	}
}

func initWebServer() *gin.Engine {

	server := gin.Default()

	// todo 添加限流

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
	store := memstore.NewStore([]byte("OxBC4Y5fWGcsGoTlQpYIr4HzDAcZTyk4kbUFK2MqSPyWKJJ9A3JrStYXOYktSb3B"),
		[]byte("PiZuMqQnEdbgsfHcYwBoduDtGbnaK7dj"))
	//store := memstore.NewStore([]byte("secret"))
	//store, _ := redis.NewStore(6, "tcp", config.Config.Redis.Addr, "", []byte("OxBC4Y5fWGcsGoTlQpYIr4HzDAcZTyk4kbUFK2MqSPyWKJJ9A3JrStYXOYktSb3B"),
	//	[]byte("PiZuMqQnEdbgsfHcYwBoduDtGbnaK7dj"))
	server.Use(sessions.Sessions("mysession", store))

	loginMiddleWareBuilder := middleware.NewLoginMiddleWareBuilder()
	server.Use(loginMiddleWareBuilder.
		IgnorePath("/users/signup").
		IgnorePath("/users/login").
		IgnorePath("/users/login_sms/code/send").
		IgnorePath("/users/login_sms").
		Builder())

	server.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, " pong")
	})

	return server
}

func initUser(db *gorm.DB, rdb redis.Cmdable) *web.UserHandler {
	userDao := dao.NewUserDao(db)
	redisUserCache := redis2.NewRedisUserCache(rdb)
	repo := repository.NewUserRepository(userDao, redisUserCache)
	userService := service.NewUserService(repo)
	codeCache := redis2.NewCodeCache(rdb)
	codeRepository := repository.NewCodeRepository(codeCache)
	smsService := memory.NewService()
	codeService := service.NewCodeService(codeRepository, smsService)
	userHandler := web.NewUserHandler(userService, codeService)
	return userHandler
}

func initRedis() redis.Cmdable {
	return redis.NewClient(
		&redis.Options{Addr: config.Config.Redis.Addr})
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	db = db.Debug()
	if err != nil {
		panic("系统初始化错误")
	}
	return db
}

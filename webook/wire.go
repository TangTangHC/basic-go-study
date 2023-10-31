//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/TangTangHC/basic-go-study/webook/internal/repository"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/cache/redis"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/dao"
	"github.com/TangTangHC/basic-go-study/webook/internal/service"
	"github.com/TangTangHC/basic-go-study/webook/internal/web"
	"github.com/TangTangHC/basic-go-study/webook/ioc"
)

func InitWebServer() *gin.Engine {
	// 传入各组件方法
	wire.Build(
		ioc.InitRedis,
		ioc.InitDB,

		dao.NewUserDao,

		redis.NewCodeCache,
		redis.NewUserCache,

		repository.NewCodeRepository,
		repository.NewUserRepository,

		ioc.InitSmsService,

		service.NewCodeService,
		service.NewUserService,

		web.NewUserHandler,

		ioc.InitMiddleWare,
		ioc.InitWebServer,
	)
	return new(gin.Engine)
}

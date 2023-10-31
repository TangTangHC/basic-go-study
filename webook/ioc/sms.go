package ioc

import (
	"github.com/TangTangHC/basic-go-study/webook/internal/service/sms"
	"github.com/TangTangHC/basic-go-study/webook/internal/service/sms/memory"
)

func InitSmsService() sms.Service {
	return memory.NewService()
}

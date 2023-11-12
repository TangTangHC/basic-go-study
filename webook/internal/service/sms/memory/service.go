package memory

import (
	"context"
	"fmt"
	"github.com/TangTangHC/basic-go-study/webook/internal/service/sms"
)

type MomorySmsService struct {
}

func NewService() sms.Service {
	return &MomorySmsService{}
}

func (s *MomorySmsService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	fmt.Printf("执行到这了")
	return nil
}

package service

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/TangTangHC/basic-go-study/webook/internal/repository"
	"github.com/TangTangHC/basic-go-study/webook/internal/service/sms"
)

const codeTplId = "1877556"

type CodeService interface {
	Send(ctx context.Context, biz string, phone string) error
	Verify(ctx context.Context, biz string, phone string, code string) (bool, error)
}

type codeService struct {
	codeRepo repository.CodeRepository
	smsSvc   sms.Service
}

func NewCodeService(codeRepo repository.CodeRepository, smsSvc sms.Service) CodeService {
	return &codeService{
		codeRepo: codeRepo,
		smsSvc:   smsSvc,
	}
}

func (svc *codeService) Send(ctx context.Context, biz string, phone string) error {
	code := svc.generateCode()
	err := svc.codeRepo.Store(ctx, biz, phone, code)
	if err != nil {
		return err
	}
	err = svc.smsSvc.Send(ctx, codeTplId, []string{code}, phone)
	if err != nil {
		// 是否需要重试
	}
	return err
}

func (svc *codeService) Verify(ctx context.Context, biz string, phone string, code string) (bool, error) {
	return svc.codeRepo.Verify(ctx, biz, phone, code)
}

func (svc *codeService) generateCode() string {
	intn := rand.Intn(1000000)
	return fmt.Sprintf("%6d", intn)
}

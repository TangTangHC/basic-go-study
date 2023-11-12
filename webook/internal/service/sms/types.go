package sms

import (
	"context"
	"errors"
)

type Service interface {
	Send(ctx context.Context, tpl string, args []string, numbers ...string) error
}

type Limiter interface {
	Limit(ctx context.Context, key string) (bool, error)
}

type LimitSmsSendV1 struct {
	Service
	limiter Limiter
}

func NewLimitSmsSendV1(limiter Limiter) *LimitSmsSend {
	return &LimitSmsSend{
		limiter: limiter,
	}
}

func (l *LimitSmsSendV1) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	limit, err := l.limiter.Limit(ctx, "")
	if err != nil {
		return err
	}
	if !limit {
		return errors.New("出发限流")
	}
	err = l.Service.Send(ctx, tpl, args, numbers...)
	return err
}

type LimitSmsSend struct {
	svc     Service
	limiter Limiter
}

func NewLimitSmsSend(service Service, limiter Limiter) *LimitSmsSend {
	return &LimitSmsSend{
		svc:     service,
		limiter: limiter,
	}
}

func (l *LimitSmsSend) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	limit, err := l.limiter.Limit(ctx, "")
	if err != nil {
		return err
	}
	if !limit {
		return errors.New("出发限流")
	}
	err = l.svc.Send(ctx, tpl, args, numbers...)
	return err
}

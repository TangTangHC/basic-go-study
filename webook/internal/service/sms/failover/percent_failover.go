package failover

import (
	"context"
	_ "embed"
	"errors"
	"github.com/TangTangHC/basic-go-study/webook/internal/service/sms"
	"github.com/redis/go-redis/v9"
	"reflect"
)

//go:embed lua/slide_window.lua
var percentSmsLua string

type PercentFailoverSmsService struct {
	smsSvcs           []sms.Service
	svcsMap           map[string]sms.Service
	cli               redis.Cmdable
	maxErrCnt         int64
	storeFailDataFunc func(string, []string, ...string) (int64, error)

	errChan chan error
}

func NewPercentFailoverSmsService(smsSvcs []sms.Service, cli redis.Cmdable,
	maxErrCnt int64, storeFailDataFunc func(string, []string, ...string) (int64, error), errChan chan error) sms.Service {
	hash := make(map[string]sms.Service, len(smsSvcs))
	for _, v := range smsSvcs {
		var name string
		of := reflect.TypeOf(v)
		if of.Kind() == reflect.Ptr {
			name = of.Elem().Name()
		} else {
			name = of.Name()
		}
		hash[name] = v
	}
	return PercentFailoverSmsService{
		smsSvcs:           smsSvcs,
		svcsMap:           hash,
		cli:               cli,
		storeFailDataFunc: storeFailDataFunc,
		maxErrCnt:         maxErrCnt,
		errChan:           errChan,
	}
}

func (p PercentFailoverSmsService) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	result, err := p.cli.ZRevRangeWithScores(ctx, p.key(), 0, 0).Result()
	if len(result) == 0 || err != nil {
		return errors.New("系统异常")
	}
	key := result[0].Member
	service, ok := p.svcsMap[key.(string)]
	if !ok {
		return errors.New("服务不可用")
	}
	err = service.Send(ctx, tpl, args, numbers...)
	if err != nil {
		go func() {
			//err := p.storeFailDataFunc(tpl, args, numbers...)()
			//if err != nil {
			//	p.errChan <- err
			//} else {
			//	p.errChan <- nil
			//}
		}()
	}
	return nil
}

func (p PercentFailoverSmsService) key() string {
	return "PercentFailoverSms"
}

func t() {
	errChan := make(chan error)
	service := NewPercentFailoverSmsService([]sms.Service{}, nil, 100, nil, errChan)
	err := service.Send(context.Background(), "", []string{})
	if err != nil {
		// 方法结束
	}

	go func() {
		select {
		case err1 := <-errChan:
			if err1.Error() == "存储失败" {
				// 处理存储失败的情况
			}
			close(errChan)
		}
	}()

}

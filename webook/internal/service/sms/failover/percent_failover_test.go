package failover

import (
	"context"
	_ "embed"
	"github.com/TangTangHC/basic-go-study/webook/internal/service/sms"
	"github.com/TangTangHC/basic-go-study/webook/internal/service/sms/memory"
	"github.com/redis/go-redis/v9"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
	"time"
)

//go:embed lua/slide_window.lua
var slideWindowLua string

func TestPercentFailoverSmsService_Send(t *testing.T) {
	var testCase = []struct {
		name     string
		mockFunc func(controller *gomock.Controller) ([]sms.Service, redis.Cmdable)

		storeFailDataFunc func(string, []string, ...string) error
		interval          time.Duration
		rate              int
	}{
		{
			name: "正常使用",
			mockFunc: func(ctl *gomock.Controller) ([]sms.Service, redis.Cmdable) {
				//s1 := smsmocks.NewMockService(ctl)
				//s1.EXPECT().Send()
				return nil, nil
			},
		},
	}
	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			//ctl := gomock.NewController(t)
			//svcs, cmdable := tc.mockFunc(ctl)
			//NewPercentFailoverSmsService(svcs, cmdable, tc.storeFailDataFunc, tc.interval, tc.rate)
			//service.Send()
		})
	}
}

func Test_lua(t *testing.T) {
	cli := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	service := memory.NewService()

	of := reflect.TypeOf(service)
	var key string
	if of.Kind() == reflect.Ptr {
		key = of.Elem().Name()
	} else {
		key = of.Name()
	}
	t.Log(key)
	b, err2 := cli.Eval(context.Background(), slideWindowLua, []string{key}, 100, 10, time.Now().Unix()).Bool()
	t.Log(err2)
	t.Log(b)

	var a interface{}
	a = "aaa"
	//fmt.Println(a.(string))
	switch a.(type) {
	case bool:
	case string:
		t.Log("=====")
	}
}

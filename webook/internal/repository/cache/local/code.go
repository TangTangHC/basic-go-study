package local

import (
	"context"
	"github.com/TangTangHC/basic-go-study/webook/internal/repository/cache/redis"
	"sync"
	"time"
)

type LocalCodeCache struct {
	hash sync.Map
	sec  int
}

type value struct {
	starTime   time.Time
	endTime    time.Time
	visitCount uint
	code       string
}

type queueValue struct {
	key     string
	endTime time.Time
}

func NewLocalCodeCache(sec int) redis.CodeCache {
	if sec <= 0 {
		sec = 60 * 5
	}
	l := &LocalCodeCache{
		sec: sec,
	}
	ticker := time.NewTicker(time.Duration(sec) * time.Second)
	stop := make(chan struct{})
	go func() {
		for {
			now := time.Now()
			select {
			case <-ticker.C:
				// 遍历reQ，删除hash里面过期的key
				l.hash.Range(func(key, value any) bool {
					if isExpired(now, value) {
						l.hash.Delete(key)
					}
					return true
				})
			case <-stop:
				ticker.Stop()
				return
			}
		}
	}()
	return l
}

func isExpired(now time.Time, v any) bool {
	value, ok := v.(*value)
	if !ok {
		return true
	}
	if now.After(value.endTime) {
		return false
	}
	return true
}

func (l *LocalCodeCache) Set(ctx context.Context, biz, phone, code string) error {
	key := biz + ":" + phone
	load, ok := l.hash.Load(key)
	now := time.Now()
	if ok {
		v := load.(*value)
		// 判断时间间隔
		if now.Before(v.starTime.Add(540 * time.Second)) {
			return redis.ErrCodeSendTooMany
		}
	}
	l.hash.Store(key, &value{
		starTime: now,
		endTime:  now.Add(600 * time.Second),
		//endTime:    now.Add(600 * time.Millisecond),
		visitCount: 3,
		code:       code,
	})
	return nil
}

func (l *LocalCodeCache) Verify(ctx context.Context, biz, phone, inputCode string) (bool, error) {
	key := biz + ":" + phone
	load, ok := l.hash.Load(key)
	if !ok {
		return false, nil
	}
	// 判断过期时间
	v := load.(*value)
	if v.endTime.Before(time.Now()) {
		l.hash.Delete(key)
		return false, nil
	}
	if v.visitCount == 0 {
		l.hash.Delete(key)
		return false, redis.ErrCodeVerifyTooManyTimes
	}
	if v.code != inputCode {
		v.visitCount--
		l.hash.Store(key, v)
		return false, nil
	}
	return true, nil
}

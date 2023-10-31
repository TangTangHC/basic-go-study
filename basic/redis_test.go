package main

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

func Test_redis(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "localhost:6379",
	})
	type User struct {
		Name string
		Age  int
	}
	u := &User{
		Name: "thc",
		Age:  19,
	}
	marshal, _ := json.Marshal(u)
	result, err := client.Set(context.Background(), "hh", marshal, 10*time.Minute).Result()
	t.Log(result)
	t.Log(err)

	//val := client.Get(context.Background(), "hh").Val()
	//t.Log("val=======", val)
	//bytes, err := client.Get(context.Background(), "hh").Bytes()
	//var u1 User
	//json.Unmarshal(bytes, &u1)
	//t.Logf("u1== %#v", u1)
	//t.Logf("u1== %+v", u1)
	//t.Log("bytes=======", string(bytes))
	//res, _ := client.Get(context.Background(), "hh").Result()
	//t.Log("res=====", res)
	//
	//name := client.Get(context.Background(), "hh").Name()
	//t.Log("name=====", name)
}

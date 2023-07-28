package basic

import (
	"fmt"
	"testing"
)

func Test_map_1(t *testing.T) {
	u := []User{
		{name: "zs"},
		{name: "ls"},
	}
	hash := make(map[string]*User)
	for _, user := range u {
		hash[user.name] = &user
	}
	for k, v := range hash {
		fmt.Printf("name: %s, user: %v\r\n", k, v)
	}
}

func Test_map_2(t *testing.T) {
	u := []*User{
		{name: "zs"},
		{name: "ls"},
	}
	hash := make(map[string]*User)
	for _, user := range u {
		hash[user.name] = user
	}
	for k, v := range hash {
		fmt.Printf("name: %s, user: %v\r\n", k, v)
	}
}

type User struct {
	name string
}

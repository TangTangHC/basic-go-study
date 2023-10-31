package main

import (
	"fmt"
	"testing"
	"time"
	"unicode/utf8"
)

func Test_func1(t *testing.T) {
	s := "您好，goland"
	size := utf8.RuneCountInString(s)
	fmt.Println(size)
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}

	fmt.Printf("%s字段最大可填写%d\n", "sig", 12)

	res, err := time.Parse("2006-01-02", "2023-19-09")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

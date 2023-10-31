package domain

import "time"

type User struct {
	Id       int64
	Email    string
	Phone    string
	Password string

	NikeName  string
	Birthday  string
	Signature string
	Ctime     time.Time
}

package domain

type User struct {
	Id       int64
	Email    string
	Password string

	NikeName  string
	Birthday  string
	Signature string
}

package repository

import (
	"fmt"

	"github.com/TangTangHC/basic-go-study/wire/repository/dao"
)

type UserRepository struct {
	ud *dao.UserDao
}

func NewUserRepository(ud *dao.UserDao) *UserRepository {
	return &UserRepository{
		ud: ud,
	}
}

func (u *UserRepository) Find() {
	u.ud.FindOne()
	fmt.Println("repository===")
}

//go:build wireinject

package wire

import (
	"github.com/TangTangHC/basic-go-study/wire/repository/dao"
	"github.com/google/wire"

	"github.com/TangTangHC/basic-go-study/wire/repository"
)

func InitRepository() *repository.UserRepository {
	wire.Build(InitDB, repository.NewUserRepository, dao.NewUserDao)
	return new(repository.UserRepository)
}

package bootstrap

import (
	urepo "github.com/buy-now-pay-later/business/users/repository"
	uucase "github.com/buy-now-pay-later/business/users/usecase"
	"github.com/buy-now-pay-later/domain/interfaces"
)

var UserUseCase interfaces.IUserUseCase

func Init() {
	userRepo := urepo.NewUserRepository()
	UserUseCase = uucase.NewUserUseCase(userRepo)
}

package usecase

import (
	"github.com/buy-now-pay-later/domain/entity"
	"github.com/buy-now-pay-later/domain/interfaces"
)

type userUseCase struct {
	userRepo interfaces.IUserRepo
}

// , txnRepo interfaces.ITxnRepo, mRepo interfaces.IMerchantRepo)

func NewUserUseCase(uRepo interfaces.IUserRepo) interfaces.IUserUseCase {
	return &userUseCase{userRepo: uRepo}
}

func (uucase *userUseCase) RegisterUser(user entity.User) {
	uucase.userRepo.AddUser(user)
}

func (uucase *userUseCase) DeleteUser(user entity.User) {
	uucase.userRepo.DeleteUser(user)
}

func (uucase *userUseCase) UpdateUser(user entity.User) error {
	return uucase.userRepo.UpdateUser(user)
}

// func (uucase *userUseCase) Transact()

func (uucase *userUseCase) PayBack(userId int, amount float64) error {
	err := uucase.userRepo.Payback(userId, amount)
	if err != nil {
		return err
	}
	return nil
}

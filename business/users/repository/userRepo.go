package repository

import (
	"errors"

	userdb "github.com/buy-now-pay-later/db/userDB"
	"github.com/buy-now-pay-later/domain/entity"
	"github.com/buy-now-pay-later/domain/interfaces"
)

type userRepo struct {
}

func NewUserRepository() interfaces.IUserRepo {
	return &userRepo{}
}

func (urepo *userRepo) AddUser(user entity.User) {
	userdb.AddUser(user)
}

func (urepo *userRepo) DeleteUser(user entity.User) {
	userdb.DeleteUser(user)
}

func (urepo *userRepo) UpdateUser(user entity.User) error {
	return userdb.UpdateUser(user)
}

func (urepo *userRepo) GetUserById(id int) (entity.User, error) {
	user := userdb.GetUserById(id)
	if user == nil {
		return entity.User{}, errors.New("no such user found")
	}
	return *user, nil
}

func (urepo *userRepo) GetAllUsersAtCreditLimit() []int {
	userIdList := userdb.GetUserAtCreditLimit()
	return userIdList
}

func (urep0 *userRepo) GetUsersWithDues() []entity.User {
	return userdb.GetUsersWithDues()
}

func (urepo *userRepo) Payback(userId int, amount float64) error {
	err := userdb.UpdateUserCreditUsed(userId, amount)
	if err != nil {
		return err
	}
	return nil
}

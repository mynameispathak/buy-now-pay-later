package interfaces

import "github.com/buy-now-pay-later/domain/entity"

type UserUseCase interface {
	RegisterUser(user entity.User) error
	// DeleteUser(user entity.User) error
	// UpdateUser(user entity.User) error
	// Transact(user entity.User, txn entity.Transaction) error
	// PayBack(userId int, amount float64) error
}

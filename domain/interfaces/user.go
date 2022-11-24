package interfaces

import "github.com/buy-now-pay-later/domain/entity"

type IUserUseCase interface {
	RegisterUser(user entity.User)
	DeleteUser(user entity.User)
	UpdateUser(user entity.User) error
	// Transact(user entity.User, txn entity.Transaction) error
	PayBack(userId int, amount float64) error
}

type IUserRepo interface {
	AddUser(user entity.User)
	DeleteUser(user entity.User)
	UpdateUser(user entity.User) error
	GetUserById(id int) (entity.User, error)
	GetAllUsersAtCreditLimit() []int
	GetUsersWithDues() []entity.User
	Payback(userId int, amount float64) error
}

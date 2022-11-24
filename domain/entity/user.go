package entity

type User struct {
	Name     string
	Id       int
	IsActive bool
	Email    string
	Account
}

type Account struct {
	CreditLimit float64
	CreditUsed  float64 // same as the amount left to be paid back by the user
}

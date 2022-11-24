package entity

type Transaction struct {
	Id         int
	UserId     int
	MerchantId int
	Amount     float64
	Discount   float64
}

package entity

type Merchant struct {
	Name        string
	Id          int
	EmailId     string
	IsActive    bool
	PromotionId string
}

type Promotion struct {
	Name     string
	Id       string
	Discount float64
}

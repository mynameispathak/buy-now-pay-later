package controller

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/buy-now-pay-later/bootstrap"
	"github.com/buy-now-pay-later/domain/entity"
	"github.com/buy-now-pay-later/helper"
)

func RegisterUser(reader *bufio.Reader) {
	fmt.Println("Enter the details of the new user in the format userName, userEmail, credtLimit")
	vals := helper.ReadVals(reader)
	if len(vals) < 3 {
		fmt.Println("invalid user format!")
		return
	}

	creditLimit, err := strconv.ParseFloat(vals[2], 64)
	if err != nil {
		fmt.Println("Invalid credit limit entered. Please try again!")
		return
	}

	user := entity.User{
		Name:  vals[0],
		Email: vals[1],
		Account: entity.Account{
			CreditLimit: creditLimit,
		},
	}

	bootstrap.UserUseCase.RegisterUser(user)
	fmt.Println("User successfully registered!")
}

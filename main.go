package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/buy-now-pay-later/bootstrap"
	userController "github.com/buy-now-pay-later/business/users/controller"
)

func main() {
	bootstrap.Init()
	Start()
}

func Start() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Buy Now Pay Later Application!")
	for {
		fmt.Println("Enter the action you want to perform:")
		fmt.Println("1. New User")
		fmt.Println("2. New Transaction")

		text, _ := reader.ReadString('\n')
		action, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
		if err != nil {
			fmt.Println("Something bad happened")
		}

		switch action {
		case 1:
			userController.RegisterUser(reader)
		default:
			fmt.Println("Invalid action choosen. Please try again!")
		}
	}
}

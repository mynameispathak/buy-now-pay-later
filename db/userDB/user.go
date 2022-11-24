package userdb

import (
	"container/list"
	"errors"
	"sync"

	"github.com/buy-now-pay-later/domain/entity"
)

var (
	users        = list.New()
	userCounter  = 0
	userIdMap    = make(map[int]*entity.User)
	addUserMutex sync.Mutex
)

func AddUser(user entity.User) {
	addUserMutex.Lock()
	{
		user.Id = userCounter
		users.PushBack(user)
		u := users.Back().Value.(entity.User)
		userIdMap[user.Id] = &u
		userCounter += 1
	}
	addUserMutex.Unlock()
}

func DeleteUser(user entity.User) {
	user.IsActive = false
}

func UpdateUser(user entity.User) error {
	if _, ok := userIdMap[user.Id]; !ok {
		return errors.New("no such user exists")
	}
	userIdMap[user.Id] = &user

	return nil
}

func GetUserById(id int) *entity.User {
	return userIdMap[id]
}

func GetUserAtCreditLimit() (userAtCreditLimitList []int) {
	userAtCreditLimitList = make([]int, 0)
	for id, user := range userIdMap {
		if user.CreditUsed == user.CreditLimit {
			userAtCreditLimitList = append(userAtCreditLimitList, id)
		}
	}
	return userAtCreditLimitList
}

func GetUsersWithDues() (userWithDuesList []entity.User) {
	userWithDuesList = make([]entity.User, 0)
	for _, user := range userIdMap {
		if user.CreditUsed > 0 {
			userWithDuesList = append(userWithDuesList, *user)
		}
	}
	return userWithDuesList
}

// Credit Payback by user
func UpdateUserCreditUsed(userId int, amount float64) error {
	if _, ok := userIdMap[userId]; !ok {
		return errors.New("no such user exists")
	}

	if userIdMap[userId].CreditUsed < amount {
		return errors.New("repay amount cannot be greater than the credit used by user")
	}

	userIdMap[userId].CreditUsed -= amount

	return nil
}

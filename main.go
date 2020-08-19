package main

import (
	"bank/accounts"
	"fmt"
)

type Account interface {
	Withdraw(value float64) string
}

func main() {
	fabioAccount := accounts.Checking{}

	fabioAccount.Deposit(2000)
	PayBill(&fabioAccount, 50)

	fmt.Println(fabioAccount.GetBalance())
}

func PayBill(account Account, value float64) {
	account.Withdraw(value)
}

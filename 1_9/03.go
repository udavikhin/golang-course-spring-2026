package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (a *BankAccount) Deposit(amount float64) {
	a.Balance += amount
}

func (a *BankAccount) Withdraw(amount float64) error {
	if a.Balance < amount {
		return errors.New("недостаточно средств")
	}

	a.Balance -= amount
	return nil
}

func main() {
	account := BankAccount{
		Owner:   "Anatoly Udavikhin",
		Balance: 4130,
	}

	account.Deposit(300)
	err := account.Withdraw(5000)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account.Balance)
}

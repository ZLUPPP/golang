package main

import (
	"fmt"
	"errors"
)


type BankAccount struct{
	Owner string
	balance float64
}

func (ba *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		ba.balance += amount
	}
} 

func (ba *BankAccount) Withdraw(amount float64) error {
	if amount > ba.balance {
		return errors.New("Not enough money")
	}
	
	if amount <= 0 {
		return errors.New("invalid amount")
	}

	ba.balance -= amount
	return nil
}

func (ba *BankAccount) Balance() float64 { return ba.balance}

func main() {
	accounts := []BankAccount{{Owner: "Beka", balance: 100}, {Owner: "Didar", balance: 200}, {Owner: "Ilya", balance: 1000}}

	for i := range accounts {
		accounts[i].Deposit(accounts[i].Balance() * 0.05)
	}

	fmt.Println(accounts)
}
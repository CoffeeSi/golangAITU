package main

import "github.com/CoffeeSi/golangAITU/assignment1/Bank"

func main() {
	account := Bank.BankAccount{
		Name:          "Yevgeniy",
		AccountNumber: 12345678,
		Money:         0,
		Transactions:  []string{},
	}
	account.ConsoleMenu()
}

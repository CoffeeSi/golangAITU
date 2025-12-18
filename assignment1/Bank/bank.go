package Bank

import (
	"fmt"
)

type BankAccount struct {
	Name          string
	AccountNumber uint32
	Money         float64
	Transactions  []string
}

func (account *BankAccount) Deposit(amount float64) {
	account.Money += amount
	trans := fmt.Sprintf("+%.2f", amount)
	account.Transactions = append(account.Transactions, trans)
	fmt.Printf("Deposit: %.2f tenge\n", amount)
}
func (account *BankAccount) Withdraw(amount float64) {
	if amount > account.Money {
		fmt.Println("You do not have enough money!")
		return
	}
	account.Money -= amount
	trans := fmt.Sprintf("-%.2f", amount)
	account.Transactions = append(account.Transactions, trans)
	fmt.Printf("Withdrawal: %.2f tenge\n", amount)
}
func (account BankAccount) GetBalance() float64 {
	return account.Money
}

func (account BankAccount) ConsoleMenu() {
	for {
		fmt.Print("\n--- Bank ---\n 1. Deposit money\n 2. Withdraw money\n 3. Display balace \n 0. Exit \n: ")
		var operation int
		fmt.Scan(&operation)
		switch operation {
		case 1:
			var amount float64
			fmt.Print("Enter the amount to deposit: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			var amount float64
			fmt.Print("Enter the amount to withdraw: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			fmt.Printf("Account balance: %.2f tenge\n", account.Money)
			fmt.Println("History:")
			for i, trans := range account.Transactions {
				fmt.Printf("%v. %s\n", i+1, trans)
			}
		case 0:
			return
		default:
			fmt.Println("Invalid number of operation!")
		}
	}
}

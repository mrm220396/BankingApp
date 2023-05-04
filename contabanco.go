package main

import "fmt"

type AccountBank struct {
	FirstName string
	LastName  string
	Balance   int
}

func (cb *AccountBank) AccountAccess() {

	var FirstName, LastName string
	fmt.Print("Type your First Name: ")
	fmt.Scan(&FirstName)
	fmt.Print("Type your last names: ")
	fmt.Scan(&LastName)

	cb.FirstName = FirstName
	cb.LastName = LastName
}

func (cb *AccountBank) Deposit() {

	var Balance int

	fmt.Print("Write an amount to deposit: R$")
	fmt.Scan(&Balance)

	cb.Balance += Balance
	fmt.Println("Deposit successfully done")

}

func (cb *AccountBank) Withdraw() (int, error) {

	var amount int

	fmt.Print("Write an amount to withdraw R$")
	fmt.Scan(&amount)

	if amount >= cb.Balance {
		err := fmt.Errorf("insufficient funds")
		return 0, err
	}

	cb.Balance -= amount

	fmt.Printf("\nWithdrawal of R$%d is successfully done\nCurrent balance R$%d\n", amount, cb.Balance)

	return cb.Balance, nil

}

func (cb *AccountBank) CheckBalance() {
	fmt.Printf("\nYour current Balance is of R$%d\n", cb.Balance)
}

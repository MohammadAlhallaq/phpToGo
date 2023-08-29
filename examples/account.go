package account

import "fmt"

type Account struct {
	Balance int
}

type AdultAccount struct {
	Account
}

type MinorAccount struct {
	Account
	Limit int
}

func (a *Account) Withdraw(amount int) {
	a.Balance -= amount
}

func (a *Account) SendStatement() {
	fmt.Println("Balance is", a.Balance)
}

func (a *MinorAccount) Withdraw(amount int) {
	if a.Account.Balance-amount < a.Limit {
		fmt.Println("Account limit exceeded")
	} else {
		a.Account.Withdraw(amount)
	}
}

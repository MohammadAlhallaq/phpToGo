package account

import (
	"fmt"
)

type Contactable interface {
	openingPhrase() string
}

func SendEmail(c Contactable, msg string) string {
	message := c.openingPhrase() + " " + msg
	return message
}

type Account struct {
	Balance int
}

type AdultAccount struct {
	Account
}

func (a AdultAccount) openingPhrase() string {
	return "this is from an adult account"
}

type MinorAccount struct {
	Account
	Limit int
}

func (a *MinorAccount) openingPhrase() string {
	return "this is from an minor account"

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

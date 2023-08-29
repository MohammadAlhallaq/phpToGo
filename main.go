package main

import (
	account "phpToGo/examples"
)

func main() {

	minor := account.MinorAccount{
		Limit:   100,
		Account: account.Account{Balance: 600},
	}

	minor.Withdraw(100)
	minor.SendStatement()
}

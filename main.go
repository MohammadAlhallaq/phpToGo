package main

import (
	"fmt"
	account "phpToGo/examples"
)

func main() {

	adult := account.MinorAccount{
		Account: account.Account{Balance: 1000},
	}

	msg := account.SendEmail(&adult, "test")

	fmt.Println(msg)
}

package main

import (
	"fmt"
	"phpToGo/methods"
)

func main() {
	amounts := methods.MoneyAmounts{500, 300}
	amounts.Add(200)
	fmt.Println(amounts.Sum())
}

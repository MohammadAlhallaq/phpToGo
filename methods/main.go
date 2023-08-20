package main

import "fmt"

type money uint
type moneyAmounts []money

func (m *moneyAmounts) add(amount money) {
	*m = append(*m, amount)
}

func (m moneyAmounts) sum() money {
	var total money
	for i := 0; i < len(m); i++ {
		total += m[i]
	}
	return total
}

func main() {
	amounts := moneyAmounts{500, 300}
	amounts.add(200)
	fmt.Println(amounts.sum())
}

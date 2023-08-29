package methods

type Money uint
type MoneyAmounts []Money

func (m MoneyAmounts) Add(amount Money) {
	m = append(m, amount)
}

func (m MoneyAmounts) Sum() Money {
	var total Money
	for i := 0; i < len(m); i++ {
		total += m[i]
	}
	return total
}

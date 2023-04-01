package generics

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjustBalance := func(currentBalance float64, trans Transaction) float64 {
		if trans.From == name {
			return currentBalance - trans.Sum
		}

		if trans.To == name {
			return currentBalance + trans.Sum
		}

		return currentBalance
	}

	return Reduce(transactions, adjustBalance, 0)
}

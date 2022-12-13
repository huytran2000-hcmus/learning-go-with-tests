package pointers

import (
	"fmt"
)

type Bitcoin int

type InsufficientFundsErr struct {
	Withdraw Bitcoin
	Funds    Bitcoin
}

func (err InsufficientFundsErr) Error() string {
	return fmt.Sprintf("Want %d bitcoins, wallet have %d bitcoins", err.Withdraw, err.Funds)
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsErr{
			Withdraw: amount,
			Funds:    w.balance,
		}
	}

	w.balance -= amount
	return nil
}

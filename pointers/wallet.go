package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) bool {
	w.balance += amount
	return true
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if (w.balance - amount) < 0 {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

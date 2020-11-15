package wallet

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds : var for insufficient btc
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin type
type Bitcoin int

// Stringer interface
type Stringer interface {
	String() string
}

// Wallet struct
type Wallet struct { 
	balance Bitcoin
}

// Deposit wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance func
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Withdraw btc
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
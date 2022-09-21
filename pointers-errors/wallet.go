package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// w (reciever method) is a copy, not the value of instance
// with * (pointer to ->) use a pointer to that value
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// by convention if one mtehod have *
// all methods will have * (pointer)
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}

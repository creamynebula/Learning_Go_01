package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("vacilo! tentando sacar mais do que tem! ")

type Bitcoin int // então 10 bitcoins == Bitcoin(10)

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) { // recebe um pointer pra wallet 'w'
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {

	if w == nil {
		return Bitcoin(0)
	}

	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

type Stringer interface { // isso quer dizer que Stringer é algum type que possui um método String(), acho
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

package main

import (
	"errors"
	"fmt"
)

type Bitcoin int // então 10 bitcoins == Bitcoin(10)

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) { // recebe um pointer pra wallet 'w'
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return errors.New("vacilo! tentando sacar mais do que tem!")
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

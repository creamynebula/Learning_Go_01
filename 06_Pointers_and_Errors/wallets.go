package main

import "fmt"

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

type Stringer interface { // isso quer dizer que Stringer é algum type que possui um método String(), acho
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

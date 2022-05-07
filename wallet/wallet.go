package wallet

import (
	"fmt"
)

type AnCoin int

var InsufficientFundError = fmt.Errorf("insufficient fund")

func (c AnCoin) String() string {
	return fmt.Sprintf("%d ANC", c)
}

type Wallet struct {
	balance AnCoin
}

func (w *Wallet) Deposit(amount AnCoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount AnCoin) error {
	if w.balance < amount {
		return fmt.Errorf("%w: withdrawal amount exceeds balance", InsufficientFundError)
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() AnCoin {
	return w.balance
}

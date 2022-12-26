package quickcash

import (
	"errors"
)

type PaytmWallet struct {
	balance float64
}

func (acc *PaytmWallet) GetBalance() float64 {
	return acc.balance
}

func (acc *PaytmWallet) GetIdentifier() string {
	return "PAYTM_WALLET"
}

func (acc *PaytmWallet) CanWithDraw(amount float64) bool {
	return acc.balance >= amount
}

func (acc *PaytmWallet) WithDraw(amount float64) error {
	if acc.CanWithDraw(amount) {
		acc.balance -= amount
		return nil
	} else {
		return errors.New("Insufficient Balance")
	}
}

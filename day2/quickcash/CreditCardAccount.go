package quickcash

import (
	"errors"
	"math"
)

type CreditCardAccount struct {
	balance   float64
	cardLimit float64
}

func (acc *CreditCardAccount) GetBalance() float64 {
	return acc.balance
}

func (acc *CreditCardAccount) CanWitdraw(amount float64) bool {
	return acc.cardLimit >= math.Abs(acc.balance-amount)
}

func (acc *CreditCardAccount) Withdraw(amount float64) error {
	if acc.CanWitdraw(amount) {
		acc.balance -= amount
		return nil
	} else {
		return errors.New("Overlimit")
	}
}

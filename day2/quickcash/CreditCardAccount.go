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

func (acc *CreditCardAccount) GetIdentifier() string {
	return "CREDITCARD_ACCOUNT"
}

func (acc *CreditCardAccount) CanWithDraw(amount float64) bool {
	return acc.cardLimit >= math.Abs(acc.balance-amount)
}

func (acc *CreditCardAccount) WithDraw(amount float64) error {
	if acc.CanWithDraw(amount) {
		acc.balance -= amount
		return nil
	} else {
		return errors.New("Overlimit")
	}
}

package quickcash

import (
	"errors"
)

type SavingsAccount struct {
	balance    float64
	minBalance float64
}

func (acc *SavingsAccount) GetBalance() float64 {
	return acc.balance
}

func (acc *SavingsAccount) GetIdentifier() string {
	return "SAVINGS_ACCOUNT"
}

func (acc *SavingsAccount) CanWitdraw(amount float64) bool {
	return (acc.balance - amount) >= acc.minBalance
}

func (acc *SavingsAccount) Withdraw(amount float64) error {
	if acc.CanWitdraw(amount) {
		acc.balance -= amount
		return nil
	} else {
		return errors.New("Insufficient Balance")
	}
}

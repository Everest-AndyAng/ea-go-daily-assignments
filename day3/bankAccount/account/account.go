package account

// TDD Bank Account app
import (
	"errors"
	"fmt"
)

type Account struct {
	balance float64
}

func (acc *Account) GetBalance() float64 {
	return acc.balance
}

func (acc *Account) Deposit(amount float64) {
	acc.balance += amount
}

func (acc *Account) Withdraw(amount float64) error {
	if acc.balance >= amount {
		acc.balance -= amount
		return nil
	}
	diff := amount - acc.balance

	return errors.New(fmt.Sprintf("Not enough balance. Current balance is %v, shortage of %v", acc.balance, diff))
}

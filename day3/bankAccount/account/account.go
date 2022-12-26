package account

// TDD Bank Account app
import (
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

	return &InsufficientBalanceError{acc.balance, diff}
}

type InsufficientBalanceError struct {
	balance        float64
	shortageAmount float64
}

func (err *InsufficientBalanceError) Error() string {
	return fmt.Sprintf("Not enough balance. Current balance is %v, shortage of %v", err.balance, err.shortageAmount)
}

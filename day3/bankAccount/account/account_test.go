package account

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetBalance(t *testing.T) {
	acc := Account{balance: 100}

	assert.Equal(t, float64(100), acc.GetBalance())
}

func TestSuccessfulDeposit(t *testing.T) {
	acc := Account{balance: 400}

	(&acc).Deposit(100)

	assert.Equal(t, float64(500), acc.GetBalance())
}

func TestSuccessfulWithdrawal(t *testing.T) {
	acc := Account{balance: 500}

	acc.Withdraw(200)

	assert.Equal(t, float64(300), acc.GetBalance())
}

func TestFailedWithdrawal(t *testing.T) {
	currentBalance := 400.0
	withdrawAmount := 500.0
	acc := Account{balance: currentBalance}

	err := acc.Withdraw(withdrawAmount)

	diff := withdrawAmount - currentBalance
	expectedError := &InsufficientBalanceError{currentBalance, diff}

	assert.Equal(t, err, expectedError)
}

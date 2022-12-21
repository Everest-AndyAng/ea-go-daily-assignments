package quickcash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnBalance(t *testing.T) {
	account := CreditCardAccount{0.0, 10000.0}
	assert.Equal(t, 0.0, account.GetBalance())
}

func TestShouldReturnTrueWhenWithdrawAmounIsWithinCardLimit(t *testing.T) {
	account := CreditCardAccount{0.0, 5000.0}
	amount := 100.0

	assert.Equal(t, true, account.CanWitdraw(amount))
}

func TestShouldReturnFalseWhenWithdrawAmountIsOverCardLimit(t *testing.T) {
	account := CreditCardAccount{-4000.0, 5000.0}
	amount := 2000.0

	assert.Equal(t, false, account.CanWitdraw(amount))
}

func TestShouldReduceCreditCardAccountBalanceWhenAmountWithdrawnIsSuccessful(t *testing.T) {
	account := CreditCardAccount{-1000.0, 5000.0}
	amount := 2000.0

	account.Withdraw(amount)

	assert.Equal(t, -3000.0, account.GetBalance())
}

func TestShouldReturnErrorWhenAmountWithdrawalIsOverLimit(t *testing.T) {
	account := CreditCardAccount{-4000.0, 5000.0}
	amount := 2000.0

	err := account.Withdraw(amount)

	if err == nil {
		t.Error("Expect overlimit error")
	}
}

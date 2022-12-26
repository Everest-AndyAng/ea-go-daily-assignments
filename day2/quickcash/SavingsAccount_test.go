package quickcash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnSavingsAccountBalance(t *testing.T) {
	account := SavingsAccount{200.0, 10.0}
	assert.Equal(t, 200.0, account.GetBalance())
}

func TestShouldReturnSavingsAccountIdentifier(t *testing.T) {
	account := SavingsAccount{100.0, 10.0}

	assert.Equal(t, "SAVINGS_ACCOUNT", account.GetIdentifier())
}
func TestShouldReturnTrueWhenThereIsEnoughBalanceInSavingsAccount(t *testing.T) {
	account := SavingsAccount{200.0, 10.0}
	amount := 100.0

	assert.Equal(t, true, account.CanWithDraw(amount))
}

func TestShouldReturnFalseWhenWithdrawAmountIsOverMinimumBalanceAllowedInSavingsAccount(t *testing.T) {
	account := SavingsAccount{100.0, 10.0}
	amount := 100.0

	assert.Equal(t, false, account.CanWithDraw(amount))
}

func TestShouldReduceSavingsAccountBalanceWhenAmountIsWithdrawalIsSuccessful(t *testing.T) {
	account := SavingsAccount{200.0, 10.0}
	amount := 100.0

	account.WithDraw(amount)

	assert.Equal(t, 100.0, account.GetBalance())
}

func TestShouldReturnErrorWhenAmountWithdrawalFailedFromSavingsAccount(t *testing.T) {
	account := SavingsAccount{100.0, 10.0}
	amount := 100.0

	err := account.WithDraw(amount)

	if err == nil {
		t.Error("Expect insufficient balance error")
	}
}

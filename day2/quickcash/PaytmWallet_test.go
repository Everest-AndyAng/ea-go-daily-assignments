package quickcash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnPaytmWalletBalance(t *testing.T) {
	account := PaytmWallet{200.0}
	assert.Equal(t, 200.0, account.GetBalance())
}

func TestShouldReturnPaytmWalletIdentifier(t *testing.T) {
	account := PaytmWallet{100.0}

	assert.Equal(t, "PAYTM_WALLET", account.GetIdentifier())
}
func TestShouldReturnTrueWhenThereIsEnoughBalanceInPaytmWallet(t *testing.T) {
	account := PaytmWallet{200.0}
	amount := 100.0

	assert.Equal(t, true, account.CanWithDraw(amount))
}

func TestShouldReturnFalseWhenThereIsNotEnoughFundsInPaytmWallet(t *testing.T) {
	account := PaytmWallet{90.0}
	amount := 100.0

	assert.Equal(t, false, account.CanWithDraw(amount))
}

func TestShouldReducePaytmWalletBalanceWhenAmountIsWithdrawalIsSuccessful(t *testing.T) {
	account := PaytmWallet{200.0}
	amount := 100.0

	account.WithDraw(amount)

	assert.Equal(t, 100.0, account.GetBalance())
}

func TestShouldReturnErrorWhenAmountWithdrawalFailedFromPaytmWallet(t *testing.T) {
	account := PaytmWallet{100.0}
	amount := 105.0

	err := account.WithDraw(amount)

	if err == nil {
		t.Error("Expect insufficient balance error")
	}
}

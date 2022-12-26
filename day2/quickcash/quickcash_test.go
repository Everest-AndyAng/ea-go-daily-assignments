package quickcash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCashFromSavingsAccount(t *testing.T) {

	fpa := &SavingsAccount{600.0, 10.0}
	fsa := &CreditCardAccount{0.0, 1000.0}

	fqc := QuickCash{
		[]Withdrawable{fpa, fsa},
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fpa.GetIdentifier(), accType)
}

func TestGetCashFromSecondaryAccount(t *testing.T) {
	fpa := &SavingsAccount{300.0, 10.0}
	fsa := &CreditCardAccount{0.0, 1000.0}

	fqc := QuickCash{
		[]Withdrawable{fpa, fsa},
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fsa.GetIdentifier(), accType)
}

func TestGetCashFromPaytmWallet(t *testing.T) {
	primaryAcc := &SavingsAccount{300.0, 10.0}
	secondaryAcc := &PaytmWallet{500.0}

	fqc := QuickCash{
		[]Withdrawable{primaryAcc, secondaryAcc},
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, secondaryAcc.GetIdentifier(), accType)
}

func TestShouldReturnWithdrawableAccountWhenThereIsEnoughFunds(t *testing.T) {
	primaryAcc := &SavingsAccount{300.0, 10.0}
	secondaryAcc := &PaytmWallet{500.0}
	tenaryAcc := &CreditCardAccount{-5000, 10000.0}

	fqc := QuickCash{
		[]Withdrawable{primaryAcc, secondaryAcc, tenaryAcc},
	}

	withdrawableAccount, _ := fqc.withdrawableAccount(2000)
	assert.Equal(t, tenaryAcc, withdrawableAccount)
}

func TestShouldReturnErrorWhenThereIsNotEnoughFundsAvaialbleForWithdrawal(t *testing.T) {
	primaryAcc := &SavingsAccount{300.0, 10.0}
	secondaryAcc := &PaytmWallet{500.0}
	tenaryAcc := &CreditCardAccount{-9000, 10000.0}

	fqc := QuickCash{
		[]Withdrawable{primaryAcc, secondaryAcc, tenaryAcc},
	}

	withdrawableAccount, err := fqc.withdrawableAccount(2000)

	assert.Nil(t, withdrawableAccount)
	if err == nil {
		t.Errorf("Should return an error when there is not enough funds")
	}
}

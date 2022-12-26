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

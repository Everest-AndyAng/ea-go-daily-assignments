package quickcash

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCashFromSavingsAccount(t *testing.T) {

	fpa := &FakePrimaryAccount{}
	fsa := &FakeSecondaryAccount{}

	fqc := QuickCash{
		fpa,
		fsa,
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fpa.GetIdentifier(), accType)
}

func TestGetCashFromSecondaryAccount(t *testing.T) {

	fpa := &FakePrimaryAccountWithZeroBalance{}
	fsa := &FakeSecondaryAccount{}

	fqc := QuickCash{
		fpa,
		fsa,
	}

	amt, accType := fqc.getCash(500)
	assert.Equal(t, float64(500), amt)
	assert.Equal(t, fsa.GetIdentifier(), accType)
}

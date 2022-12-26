package quickcash

import "errors"

// Assuming the order of accounts is always ordered
type QuickCash struct {
	Accounts []Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string) {
	withdrawableAccount, _ := qc.withdrawableAccount(amount)

	withdrawableAccount.WithDraw(amount)
	return amount, withdrawableAccount.GetIdentifier()
}

func (qc *QuickCash) withdrawableAccount(amount float64) (Withdrawable, error) {
	for _, account := range qc.Accounts {
		if account.GetIdentifier() == "SAVINGS_ACCOUNT" && account.CanWithDraw(amount) {
			return account, nil
		}
		if account.GetIdentifier() == "CREDITCARD_ACCOUNT" && account.CanWithDraw(amount) {
			return account, nil
		}
		if account.GetIdentifier() == "PAYTM_WALLET" && account.CanWithDraw(amount) {
			return account, nil
		}
	}
	return nil, errors.New("No funds avalable from all accounts")
}

package quickcash

type QuickCash struct {
	Accounts []Withdrawable
}

func (qc *QuickCash) getCash(amount float64) (float64, string) {
	withdrawableAccount := withdrawableAccount(qc, amount)

	withdrawableAccount.WithDraw(amount)
	return amount, withdrawableAccount.GetIdentifier()
}

func withdrawableAccount(qc *QuickCash, amount float64) Withdrawable {
	for _, account := range qc.Accounts {
		if account.GetIdentifier() == "SAVINGS_ACCOUNT" && account.CanWithDraw(amount) {
			return account
		}
		if account.GetIdentifier() == "CREDITCARD_ACCOUNT" && account.CanWithDraw(amount) {
			return account
		}
	}
	return nil
}

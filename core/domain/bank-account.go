package domain

type BankAccount struct {
	AccountNumber string
	Balance       float64
}

func (bankAccount *BankAccount) Debit(amount float64) {
	bankAccount.Balance -= amount
}

func (bankAccount *BankAccount) Credit(amount float64) {
	bankAccount.Balance += amount
}

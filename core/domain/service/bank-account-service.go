package service

import (
	"fullcycle91/core/domain"
	"fullcycle91/core/domain/repository"
)

type BankAccountService struct{}

type TransferBetweenAccountsStruct struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func Create(bankAccountRepository repository.BankAccountRepositoryInterface, accountNumber string) {
	bankAccount := domain.BankAccount{
		AccountNumber: accountNumber,
		Balance:       1000,
	}

	bankAccountRepository.Insert(bankAccount)
}

func Transfer(bankAccountRepository repository.BankAccountRepositoryInterface, payload TransferBetweenAccountsStruct) error {
	from := payload.From
	firstAccount := domain.BankAccount{
		AccountNumber: from,
		Balance:       0,
	}

	to := payload.To
	secondAccount := domain.BankAccount{
		AccountNumber: to,
		Balance:       0,
	}

	amount := payload.Amount

	var accounts []domain.BankAccount
	accounts = append(accounts, firstAccount)

	if firstAccount.AccountNumber == secondAccount.AccountNumber {
		accounts = bankAccountRepository.Where("account_number", accounts)
		firstAccount = accounts[0]

		firstAccount.Credit(amount)

		err := bankAccountRepository.Update("balance", firstAccount.Balance, "account_number", firstAccount.AccountNumber)
		if err != nil {
			return err
		}

		return nil
	}

	accounts = append(accounts, secondAccount)

	accounts = bankAccountRepository.Where("account_number", accounts)
	firstAccount = accounts[0]
	secondAccount = accounts[1]

	firstAccount.Debit(amount)
	secondAccount.Credit(amount)

	err := bankAccountRepository.Update("balance", firstAccount.Balance, "account_number", firstAccount.AccountNumber)

	if err != nil {
		return err
	}

	err = bankAccountRepository.Update("balance", secondAccount.Balance, "account_number", secondAccount.AccountNumber)
	if err != nil {
		return err
	}

	return nil
}

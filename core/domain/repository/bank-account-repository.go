package repository

import (
	"fullcycle91/core/domain"
)

type BankAccountRepositoryInterface interface {
	Insert(bankAccount domain.BankAccount) domain.BankAccount
	Where(columnName string, values []domain.BankAccount) []domain.BankAccount
	Update(columnNameSet string, setValue any, columnNameWhere string, columnValueWhere any) error
}

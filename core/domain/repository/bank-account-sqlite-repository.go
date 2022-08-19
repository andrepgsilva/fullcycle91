package repository

import (
	"database/sql"
	"fullcycle91/core/domain"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type BankAccountSqliteRepository struct{}

func (b BankAccountSqliteRepository) Insert(bankAccount domain.BankAccount) domain.BankAccount {
	db, err := sql.Open("sqlite3", "database/data.db")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO bank_accounts(account_number, balance) values(?,?)")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(bankAccount.AccountNumber, bankAccount.Balance)

	if err != nil {
		panic(err)
	}

	db.Close()
	return bankAccount
}

func (b BankAccountSqliteRepository) Where(columnName string, values []domain.BankAccount) []domain.BankAccount {
	db, err := sql.Open("sqlite3", "database/data.db")
	if err != nil {
		panic(err)
	}

	var queryBuilder string = "SELECT * FROM bank_accounts WHERE "

	if len(values) == 1 {
		queryBuilder += columnName + "=?"
		var id string

		row := db.QueryRow(queryBuilder, values[0].AccountNumber)

		if err := row.Scan(&id, &values[0].AccountNumber, &values[0].Balance); err != nil {
			log.Fatal(err)
		}

		return values
	}

	var accounts []domain.BankAccount
	var queryArgs []any

	for i := 0; i < len(values); i++ {
		queryArgs = append(queryArgs, values[i].AccountNumber)
		queryBuilder += columnName + "=?"

		if i == len(values)-1 {
			break
		}

		queryBuilder += " OR "
	}

	rows, err := db.Query(queryBuilder, queryArgs...)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id string
		account := domain.BankAccount{}

		if err := rows.Scan(&id, &account.AccountNumber, &account.Balance); err != nil {
			log.Fatal(err)
		}

		accounts = append(accounts, account)
	}

	rows.Close()
	db.Close()

	return accounts
}

func (b BankAccountSqliteRepository) Update(columnNameSet string, setValue any, columnNameWhere string, columnValueWhere any) error {
	db, err := sql.Open("sqlite3", "database/data.db")

	if err != nil {
		panic(err)
	}

	queryBuilder := "UPDATE bank_accounts set "
	queryBuilder += columnNameSet + "=? WHERE " + columnNameWhere + "=?"

	stmt, err := db.Prepare(queryBuilder)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(setValue, columnValueWhere)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}

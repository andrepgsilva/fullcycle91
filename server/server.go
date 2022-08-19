package server

import (
	"encoding/json"
	"fmt"
	"fullcycle91/core/domain/repository"
	"fullcycle91/core/domain/service"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
)

type JsonRequestCreateBankAccount struct {
	AccountNumber string `json:"account_number"`
}

var bankAccountRepository = repository.BankAccountSqliteRepository{}

func createBankAccountLink(w http.ResponseWriter, r *http.Request) {
	jsonRequest := JsonRequestCreateBankAccount{}

	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "An Error Ocurred")
	}

	err = json.Unmarshal(reqBody, &jsonRequest)

	if err != nil {
		panic(err)
	}

	service.Create(bankAccountRepository, jsonRequest.AccountNumber)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&jsonRequest)
}

func transferBetweenBankAccountsLink(w http.ResponseWriter, r *http.Request) {
	jsonRequest := service.TransferBetweenAccountsStruct{}

	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "An Error Ocurred")
	}

	err = json.Unmarshal(reqBody, &jsonRequest)

	if err != nil {
		panic(err)
	}

	service.Transfer(bankAccountRepository, jsonRequest)

	w.WriteHeader(http.StatusOK)
}

func Init() {
	fmt.Printf("Starting server at port 8000\n")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/bank-accounts", createBankAccountLink)
	router.HandleFunc("/bank-accounts/transfer", transferBetweenBankAccountsLink)
	log.Fatal(http.ListenAndServe(":8000", router))
}

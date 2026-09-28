package main

import "errors"

type Transaction struct {
	id          int
	amount      float64
	category    string
	description string
	date        string
}

func NewTransaction(amount float64, category string, description string, date string) Transaction {
	return Transaction{len(allTransaction) + 1, amount, category, description, date}
}

func AddTransaction(tx Transaction) error {
	if tx.amount < 0 {
		return errors.New("Сумма транзакции должна быть больше нуля")
	}
	allTransaction = append(allTransaction, tx)
	return nil
}

func ListTransactions() []Transaction {
	return allTransaction
}

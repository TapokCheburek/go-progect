package internal

import (
	"encoding/json"
	"io"
)

var allTransaction []Transaction
var allBudget = map[string]Budget{
	"Еда":      {category: "Еда", limit: 15000},
	"Переводы": {category: "Переводы", limit: 1200},
}

type Transaction struct {
	id          int
	amount      float64
	category    string
	description string
	date        string
}

type Budget struct {
	category string
	limit    float64
}

func NewTransaction(amount float64, category string, description string, date string) Transaction {
	return Transaction{len(allTransaction) + 1, amount, category, description, date}
}

func NewBudget(category string, limit float64) Budget {
	return Budget{category, limit}
}

func AddTransaction(tx Transaction) error {
	if tx.amount < 0 {
		return errNegativeAmount
	}

	if sumTransactionCategory(tx)+tx.amount > allBudget[tx.category].limit {
		return errBudget

	}
	allTransaction = append(allTransaction, tx)
	return nil
}

func sumTransactionCategory(tx Transaction) float64 {
	sum := 0.0
	for _, t := range allTransaction {
		if t.category == tx.category {
			sum += t.amount
		}
	}
	return sum
}

func SetBudget(b Budget) {
	allBudget[b.category] = b
}

func ListTransactions() []Transaction {
	return allTransaction
}

func ListBudget() map[string]Budget {
	return allBudget
}

func LoadBudgets(r io.Reader) error {
	var items []struct {
		Category string  `json:"category"`
		Limit    float64 `json:"limit"`
	}

	if err := json.NewDecoder(r).Decode(&items); err != nil {
		return errReadJSON
	}

	for _, it := range items {
		SetBudget(NewBudget(it.Category, it.Limit))
	}
	return nil
}

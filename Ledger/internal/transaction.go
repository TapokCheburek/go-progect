package internal

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

func (t Transaction) Validate() error {
	if t.amount <= 0 {
		return errNegativeLimit
	} else if t.category == "" {
		return errNullCategory
	}
	return nil
}

func NewTransaction(amount float64, category string, description string, date string) Transaction {
	return Transaction{len(allTransaction) + 1, amount, category, description, date}
}

func AddTransaction(tx Transaction) error {
	err := tx.Validate()
	if err != nil {
		return err
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

func ListTransactions() []Transaction {
	return allTransaction
}

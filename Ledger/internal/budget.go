package internal

import (
	"encoding/json"
	"io"
)

type Budget struct {
	category string
	limit    float64
}

func NewBudget(category string, limit float64) Budget {
	return Budget{category, limit}
}

func (b Budget) Validate() error {
	if b.limit <= 0 {
		return errNegativeAmount
	} else if b.category == "" {
		return errNullCategory
	}
	return nil
}

func SetBudget(b Budget) error {
	err := b.Validate()
	if err != nil {
		return err
	}
	allBudget[b.category] = b
	return nil
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

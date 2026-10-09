package main

import (
	"Ladger/internal"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("“Ledger service started”")

	f, err := os.Open("budgets.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := internal.LoadBudgets(bufio.NewReader(f)); err != nil {
		panic(err)
	}
	fmt.Println(internal.ListBudget())

	err = internal.SetBudget(internal.NewBudget("Переводы", 2000))
	err = internal.SetBudget(internal.NewBudget("Транспорт", -5000))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Добавление транзакции не превышающей лимит")
	tx := internal.NewTransaction(1200, "Переводы", "Какое-то описание", "11.09.2034")
	err = internal.AddTransaction(tx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Добавление транзакции превышающей лимит")
	tx = internal.NewTransaction(1000, "Переводы", "Домашние яблоки", "02.09.2034")
	err = internal.AddTransaction(tx)
	if err != nil {
		fmt.Println(err)
	}
	transactions := internal.ListTransactions()
	for _, tx := range transactions {
		fmt.Println(tx)
	}

	b := internal.NewBudget("Транспорт", -5000)
	err = internal.CheckValid(b)
	if err != nil {
		fmt.Println(err)
	}
	t := internal.NewTransaction(-10.56, "Покупки", "Описание покупки", "09.10.2026")
	err = internal.CheckValid(t)
	if err != nil {
		fmt.Println(err)
	}
}

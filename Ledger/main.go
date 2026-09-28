package main

import "fmt"

var allTransaction []Transaction

func main() {
	fmt.Println("“Ledger service started”")
	tx := NewTransaction(2.2, "Категория 1", "Какое-то описание", "11.09.2034")
	tx2 := NewTransaction(2.45, "Яблоки", "Домашние яблоки", "02.09.2034")
	err := AddTransaction(tx)
	if err != nil {
		fmt.Println(err)
	}
	err = AddTransaction(tx2)
	if err != nil {
		fmt.Println(err)
	}
	transactions := ListTransactions()
	for _, tx := range transactions {
		fmt.Println(tx)
	}
}

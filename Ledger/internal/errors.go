package internal

import "errors"

var errNegativeAmount = errors.New("Сумма транзакции должна быть больше нуля")

var errBudget = errors.New("Превышен бюджет по категории")

var errReadJSON = errors.New("Ошибка чтения JSON")

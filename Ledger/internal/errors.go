package internal

import "errors"

var errNegativeAmount = errors.New("Сумма транзакции должна быть больше нуля")

var errNullCategory = errors.New("Категория не должна быть пустой строкой")

var errBudget = errors.New("Превышен бюджет по категории")

var errNegativeLimit = errors.New("Лимит должен быть больше нуля")

var errReadJSON = errors.New("Ошибка чтения JSON")

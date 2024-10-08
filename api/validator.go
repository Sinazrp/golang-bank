package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/sinazrp/golang-bank/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)

	}
	return false
}

var validAmount validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if amount, ok := fieldLevel.Field().Interface().(int64); ok {
		return amount > 0
	}
	return false
}

var validAccountID validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if accountID, ok := fieldLevel.Field().Interface().(int64); ok {
		return accountID > 0
	}
	return false
}

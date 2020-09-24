package money

import "fmt"

type Money struct {
	Amount   Amount   `json:"amount"`
	Currency Currency `json:"currency"`
}

func (m *Money) Plus(p Money) (Money, error) {
	if m.Currency.Symbol != p.Currency.Symbol {
		return Money{}, fmt.Errorf("could not perform money plus")
	}
	return Money{
		Amount:   m.Amount + p.Amount,
		Currency: m.Currency,
	}, nil
}

func (m *Money) Minus(p Money) (Money, error) {
	if m.Currency.Symbol != p.Currency.Symbol {
		return Money{}, fmt.Errorf("could not perform money minus")
	}
	return Money{
		Amount:   m.Amount - p.Amount,
		Currency: m.Currency,
	}, nil
}

type Amount int64

type Currency struct {
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

func DefaultCurrency() Currency {
	return Currency{
		Symbol:      "DEF",
		Description: "Default currency",
	}
}

func AmountAbs (m Money) Money {
	if m.Amount < Amount(0) {
		return Money{
			Amount:   -m.Amount,
			Currency: m.Currency,
		}
	}
	return m
}
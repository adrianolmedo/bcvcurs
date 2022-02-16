package main

import (
	"errors"
)

var ErrCurrencyNotFound = errors.New("currency not found")

type Money struct {
	Value  float64 `json:"value"`
	Iso    string  `json:"iso"`
	Symbol string  `json:"symbol"`
}

//type Currencies []Money

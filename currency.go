package main

import (
	"errors"
)

var ErrGettingData = errors.New("error getting data")

type Currency struct {
	Value  float64 `json:"value"`
	Iso    string  `json:"iso"`
	Symbol string  `json:"symbol"`
	ID     string  `json:"-"`
}

type Currencies struct {
	Euro   Currency `json:"euro"`
	Yuan   Currency `json:"yuan"`
	Lira   Currency `json:"lira"`
	Ruble  Currency `json:"ruble"`
	Dollar Currency `json:"dollar"`
}

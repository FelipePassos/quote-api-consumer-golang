package models

type dataCurrence struct {
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}

type ApiCurrence struct {
	USD_BRL dataCurrence `json:"USD_BRL"`
}

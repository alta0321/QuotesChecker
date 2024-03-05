package main

import "time"

type Quote struct {
	Id        string `gorm:"primaryKey"`
	Currency  string
	Price     float64
	CreatedAt time.Time
}

type ExchangeRateResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
	Date  string             `json:"date"`
}

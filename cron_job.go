package main

import (
	"github.com/go-co-op/gocron"
	"strings"
	"time"
)

func runCronJob() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(30).Seconds().Do(func() {
		UpdateDataBaseQuote()
	})
	s.StartBlocking()
}

func UpdateDataBaseQuote() {
	currencies := []string{
		"EUR/USD",
		"EUR/MXN",
		"USD/EUR",
		"USD/MXN",
		"MXN/USD",
		"MXN/EUR",
	}
	for _, currency := range currencies {
		sliceOfCurrencies := strings.Split(currency, "/")
		price, _ := getExchangeRate(sliceOfCurrencies[0], sliceOfCurrencies[1])
		id := generateUniqueNumber()
		quote := Quote{Currency: currency, Id: id, Price: price.Rates[sliceOfCurrencies[1]], CreatedAt: time.Now()}
		DB.Create(&quote)
	}

}

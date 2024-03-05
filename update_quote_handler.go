package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func UpdateQuote(w http.ResponseWriter, r *http.Request) {
	log.Println("Request:", r)
	currency := r.URL.Query().Get("currency")
	if currency == "" {
		responseWithError(w, http.StatusBadRequest, "Missing 'currency' query parameter")
		return
	}
	currencies := strings.Split(currency, "/")

	ch := make(chan string)
	go func() {
		price, _ := getExchangeRate(currencies[0], currencies[1])
		id := generateUniqueNumber()
		ch <- id
		quote := Quote{Currency: currency, Id: id, Price: price.Rates[currencies[1]], CreatedAt: time.Now()}
		DB.Create(&quote)
	}()
	identification := <-ch

	responseWithJSON(w, http.StatusOK, map[string]string{"id": identification})
}

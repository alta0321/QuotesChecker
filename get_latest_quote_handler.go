package main

import (
	"log"
	"net/http"
)

func GetLatestQuote(w http.ResponseWriter, r *http.Request) {
	log.Println("Request:", r)
	currency := r.URL.Query().Get("currency")
	if currency == "" {
		responseWithError(w, http.StatusBadRequest, "Missing 'currency' query parameter")
		return
	}

	var quote Quote
	DB.Where("currency = ?", currency).Order("created_at desc").First(&quote)
	responseWithJSON(w, http.StatusOK, quote)
}

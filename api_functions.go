package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
)

func getExchangeRate(fromCurrency, toCurrency string) (*ExchangeRateResponse, error) {
	url := fmt.Sprintf("https://api.frankfurter.app/latest?from=%s&to=%s", fromCurrency, toCurrency)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data ExchangeRateResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func generateUniqueNumber() string {
	n, err := rand.Int(rand.Reader, new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil))
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", n)
}

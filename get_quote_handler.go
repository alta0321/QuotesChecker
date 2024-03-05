package main

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func GetQuote(w http.ResponseWriter, r *http.Request) {
	log.Println("Request:", r)
	id := chi.URLParam(r, "id")

	var quote Quote
	DB.Where("id = ?", id).First(&quote)

	err := DB.Where("id = ?", id).First(&quote).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		responseWithJSON(w, http.StatusNotFound, "Quote not found")
	} else {
		responseWithJSON(w, http.StatusOK, quote)

	}
}

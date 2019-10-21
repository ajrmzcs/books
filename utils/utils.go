package utils

import (
	"encoding/json"
	"github.com/ajrmzcs/books/models"
	"net/http"
)

func SendErrors(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	_ = json.NewEncoder(w).Encode(data)
}
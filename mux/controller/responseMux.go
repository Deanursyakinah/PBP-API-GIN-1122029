package controller

import (
	"encoding/json"
	"net/http"

	m "mux/model"
)

func SendErrorResponseMux(w http.ResponseWriter, kode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.ErrorResponse
	response.Status = kode //400 bad req, 404 not found, 500 internal server error, 401 unauthorized
	response.Message = message

	json.NewEncoder(w).Encode(response)
}

func SendSuccesResponseMux(w http.ResponseWriter, kode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	var response m.SuccessResponse
	response.Status = kode
	response.Message = message
	json.NewEncoder(w).Encode(response)
}

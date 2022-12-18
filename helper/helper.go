package helper

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse creates success response for the http handler
func SuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	success := struct {
		Status  int         `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonData, _ := json.Marshal(success)
	w.Write(jsonData)
}

// FailedResponse creates error response for the http handler
func FailedResponse(w http.ResponseWriter, status int, err error) {
	failed := struct {
		Status int    `json:"status"`
		Error  string `json:"error"`
	}{
		Status: status,
		Error:  err.Error(),
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	jsonData, _ := json.Marshal(failed)
	w.Write(jsonData)
}

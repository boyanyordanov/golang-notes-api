package main

import (
	"net/http"
)

// JsonSerializable: Interface describing structures that can be serialized to JSON for the API
type JsonSerializable interface {
	json() ([]byte, error)
}

// WriteJson: Helper for sending out a properly formatted JSON response
func WriteJson(w http.ResponseWriter, payload JsonSerializable, statusCode int, logger FileLogger) {
	jsonPayload, err := payload.json()

	if err != nil {
		logger.Error(err.Error())
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonPayload)
}

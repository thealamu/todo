package respond

import (
	"encoding/json"
	"net/http"
)

// JSON responds with data in JSON format
func JSON(w http.ResponseWriter, data interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	coded, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.WriteHeader(statusCode)
	w.Write(coded)
	return nil
}

// Error responds with error in plain text
func Error(w http.ResponseWriter, err error, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(err.Error()))
}

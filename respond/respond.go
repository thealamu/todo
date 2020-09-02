package respond

import (
	"encoding/json"
	"net/http"
)

// JSON responds with data in JSON format
func JSON(w http.ResponseWriter, data interface{}) error {
	coded, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Write(coded)
	return nil
}

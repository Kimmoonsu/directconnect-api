package response

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, httpCode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)

	json.NewEncoder(w).Encode(data)
}

package response

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, responseData any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(responseData)

	if err != nil {
		panic(err)
	}
}

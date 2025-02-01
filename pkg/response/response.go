package response

import (
	"encoding/json"
	"net/http"
	"poymanov/todo/pkg/helpers"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func Json(w http.ResponseWriter, responseData any, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(responseData)

	if err != nil {
		panic(err)
	}
}

func JsonError(w http.ResponseWriter, errorForResponse error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	er := ErrorResponse{Message: helpers.FirstToUpper(errorForResponse.Error())}

	err := json.NewEncoder(w).Encode(er)

	if err != nil {
		panic(err)
	}
}

func NoContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

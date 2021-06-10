package http_utils

import (
	"encoding/json"
	"github.com/alvaro259818/bookstore-utils-go/rest_errors"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{})  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err rest_errors.RestError)  {
	RespondJson(w, err.Status(), err)
}

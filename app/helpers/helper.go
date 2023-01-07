package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ApiResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ApiResponseError(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		ApiResponse(w, statusCode, struct {
			Error string `json:"error"`
		}{
			err.Error(),
		})
		return
	}
	ApiResponse(w, http.StatusBadRequest, nil)
}

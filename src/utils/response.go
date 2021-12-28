package utils

import (
	"encoding/json"
	"net/http"

	"github.com/hveda/todo/src/types"
)

func JsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		json.NewEncoder(w).Encode(
			types.Response{
				Message: "Could not parse response",
			},
		)
	}
}
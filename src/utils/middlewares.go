package utils

import (
	"fmt"
	"net/http"

	"github.com/hveda/todo/src/types"
)

func HandleMethod(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			JsonResponse(w, types.Response{Message: fmt.Sprintf("Method [%s] not allowed", r.Method)})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func HandleGet(next http.Handler) http.Handler {
	return HandleMethod("GET", next)
}

func HandlePost(next http.Handler) http.Handler {
	return HandleMethod("POST", next)
}

func HandleDelete(next http.Handler) http.Handler {
	return HandleMethod("DELETE", next)
}

func HandlePut(next http.Handler) http.Handler {
	return HandleMethod("PUT", next)
}

func HandlePatch(next http.Handler) http.Handler {
	return HandleMethod("PATCH", next)
}
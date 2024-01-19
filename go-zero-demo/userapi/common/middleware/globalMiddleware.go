package middleware

import (
	"fmt"
	"net/http"
)

type GlobalTestMiddleware struct {
}

func NewGlobalTestMiddleware() *GlobalTestMiddleware {
	return &GlobalTestMiddleware{}
}

func (m *GlobalTestMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		fmt.Println("global testmiddleware begin")
		// Passthrough to next handler if need
		next(w, r)

		fmt.Println("global testmiddleware end")
	}
}

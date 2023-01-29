package handlers

import (
	"log"
	"net/http"
)

type middleware struct{}

func NewMiddlewareHandlers() *middleware {
	return &middleware{}
}

func (m *middleware) PanicRecover(http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Panic")
			}
		}()
	})
}

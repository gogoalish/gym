package server

import (
	"net/http"
	"time"

	"project/config"
)

func NewServer(mux http.Handler) *http.Server {
	return &http.Server{
		Addr:         config.C.Server.Port,
		Handler:      mux,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 5,
	}
}

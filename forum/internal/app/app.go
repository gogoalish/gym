package app

import (
	"database/sql"
	"net/http"
	"time"

	"project/internal/routes"
)

func Run() error {
	// fun db
	db := sql.DB{}
	mux := http.NewServeMux()
	routes.InitRoutes(&db, mux)
	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 5,
	}
	return server.ListenAndServe()
}

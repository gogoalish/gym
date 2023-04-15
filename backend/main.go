package main

import (
	"database/sql"
	"log"
	"net/http"

	"gym/server"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	const port string = ":3333"
	DB, _ := sql.Open("sqlite3", "gym.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	defer DB.Close()
	repo := &server.Repo{DB}
	err := server.CreateTable(repo)
	if err != nil {
		log.Fatal(err)
	}
	service := &server.Service{repo}
	handler := &server.Handler{service}
	router := mux.NewRouter()
	server.SetupRoutes(handler, router)

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}

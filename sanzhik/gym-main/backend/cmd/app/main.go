package main

import (
	"fmt"
	"log"
	"net/http"

	"gym/server"
	"gym/server/db/migrations"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mux := http.NewServeMux()
	db, err := migrations.Connect("./server/db/gym.db")
	if err != nil {
		log.Fatalln(err)
	}
	rp := server.NewRepository(db)
	svc := server.NewService(rp)
	handler := server.NewHandler(svc)
	handler.InitRoutes(mux)

	fmt.Println("OOKKK")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln(err)
	}
}

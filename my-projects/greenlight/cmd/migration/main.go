package main

import (
	"fmt"
	"log"
	"os"

	"greenlight/pkg/postgres"
)

func main() {
	if len(os.Args) <= 1 || len(os.Args) >= 3 {
		log.Fatal("Usage: go run change_db.go <argument>")
	}
	flag := os.Args[1]
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "greenlight")

	db, err := postgres.Connect(connectionString)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	switch flag {
	case "up":
		if err := postgres.CreateTable(db, "./db/migrations/up"); err != nil {
			log.Fatalln(err)
		}
		log.Println("Successfully Created")
	case "down":
		if err := postgres.DownDB(db, "./db/migrations/down"); err != nil {
			log.Fatalln(err)
		}
		log.Println("Successfully Dropped")
	default:
		log.Fatalf("%s: unknown flag. Use: 'up' or 'down'", flag)
	}
}

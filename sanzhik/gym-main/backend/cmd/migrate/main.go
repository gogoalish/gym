package main

import (
	"log"
	"os"

	"gym/server/db/migrations"
)

func main() {
	if len(os.Args) <= 1 || len(os.Args) >= 3 {
		log.Fatal("Usage: go run change_db.go <argument>")
	}
	flag := os.Args[1]
	db, err := migrations.Connect("./server/db/gym.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	switch flag {
	case "up":
		if err := migrations.CreateTable(db, "./server/db/migrations"); err != nil {
			log.Fatalln(err)
		}
		log.Println("Successful")
	case "down":
		if err := migrations.DropAllDB(db); err != nil {
			log.Fatalln(err)
		}
		log.Println("Successful")
	default:
		log.Fatalf("%s: unknown flag. Use: 'create' or 'drop'", flag)
	}
}

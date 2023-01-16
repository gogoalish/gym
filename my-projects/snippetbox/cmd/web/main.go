package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

// func main() {
// 	addr := flag.String("addr", ":4000", "HTTP network address")
// 	flag.Parse()

// 	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
// 	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

// 	app := &application{
// 		errorLog: errorLog,
// 		infoLog:  infoLog,
// 	}

// 	srv := &http.Server{
// 		Addr:     *addr,
// 		ErrorLog: errorLog,
// 		Handler:  app.routes(),
// 	}

// 	infoLog.Println("Starting server on http://localhost" + *addr)
// 	// err := http.ListenAndServe(*addr, mux)
// 	// errorLog.Fatal(err)

//		err := srv.ListenAndServe()
//		errorLog.Fatal(err)
//	}
func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)
}

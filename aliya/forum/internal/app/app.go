package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"project/config"
	"project/internal/handlers"
	"project/internal/routes"
	"project/internal/server"
	"project/pkg/sqlite"
)

var chanErr = make(chan error, 1)

func Run() error {
	log.Println("init mux")
	mux := http.NewServeMux()

	log.Println("db connection")
	db, err := sqlite.Connect(config.C.Path.DB)
	if err != nil {
		return err
	}

	log.Println("read template")
	if err := handlers.ReadTemplate(config.C.Path.Template + "*.html"); err != nil {
		return err
	}

	log.Println("init routes")
	routes.InitRoutes(db, mux)

	log.Println("init server")
	server := server.NewServer(mux)

	log.Printf("Starting listener on http://localhost%s", config.C.Server.Port)
	go func() {
		chanErr <- server.ListenAndServe()
	}()

	return wait()
}

func wait() error {
	syscalCh := make(chan os.Signal, 1)
	signal.Notify(syscalCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-syscalCh:
		fmt.Println()
		log.Printf("Stop server...\n")
		return nil
	case err := <-chanErr:
		return err
	}
}

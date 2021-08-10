package main

import (
	"context"
	"fmt"
	"github.com/karakanb/plans-service-grpc-poc/pkg/handler"
	"github.com/karakanb/plans-service-grpc-poc/pkg/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

const (
	HTTP_PORT = 8080
)

func registerRoutes(server *handler.HttpServer) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/plans", server.CreatePlan).Methods("POST")
	r.HandleFunc("/plans/{id}", server.GetPlan).Methods("POST")
	return r
}

func main() {
	plansRepo := repository.NewPlansRepository()
	server := handler.NewHttpServer(handler.NewHandler(plansRepo))
	router := registerRoutes(server)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", HTTP_PORT),
		Handler: router,
	}

	serverErrs := make(chan error, 1)
	go func() {
		log.Printf("running the http server on port: %d\n", HTTP_PORT)
		serverErrs <- httpServer.ListenAndServe()
	}()

	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGTERM, os.Interrupt)

	select {
	case err := <-serverErrs:
		log.Fatalf("failed to start service: %v\n", err)
	case <-exit:
		log.Println("shutting down server")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(ctx); err != nil {
			log.Fatalf("server shutdown did not complete: %v\n", err)
		}
	}
}

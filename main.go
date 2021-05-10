package main

import (
	"Ex5Validation/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "product-api ", log.LstdFlags)

	productsHandler := handlers.NewProducts(logger)

	serveMux := mux.NewRouter()

	getRouter := serveMux.Methods("GET").Subrouter()
	getRouter.HandleFunc("/", productsHandler.GetProducts)

	putRouter := serveMux.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productsHandler.PutProducts)
	putRouter.Use(productsHandler.MiddlewareProductValidation)

	postRouter := serveMux.Methods("POST").Subrouter()
	postRouter.HandleFunc("/", productsHandler.PostProduct)
	postRouter.Use(productsHandler.MiddlewareProductValidation)

	server := &http.Server{
		Addr:         ":5000",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt)

	sig := <-signalChannel
	logger.Println("Received signal to terminate. Shutting Down.", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(tc)
}

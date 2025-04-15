package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/varik-08/go-metrics/internal"
)

const Host = "0.0.0.0"
const Port = "8080"

func main() {
	memStorage := internal.NewMemStorage()

	router := mux.NewRouter()
	router.HandleFunc("/update/{type}/{name}/{value}",
		func(writer http.ResponseWriter, request *http.Request) {
			addMetricHandler(writer, request, memStorage)
		}).Methods(http.MethodPost)

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}).Handler(router)

	server := &http.Server{
		Addr:         Host + ":" + Port,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Starting server on " + Host + ":" + Port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

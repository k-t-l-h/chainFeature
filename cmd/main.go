package main

import (
	chainDelivery "MarkovChain/internal/chain/delivery"
	chainUsecase "MarkovChain/internal/chain/usecase"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Print("starting")
	r := mux.NewRouter()

	cu := chainUsecase.NewUsecase()
	ch := chainDelivery.NewHandler(cu)

	r.HandleFunc("/news", ch.GetMessages).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

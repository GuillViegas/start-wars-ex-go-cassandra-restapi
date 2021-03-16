package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/GuillViegas/start-wars-ex-go-cassandra-restapi/planet"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/planet", GetUniverse).Methods("GET")
	router.HandleFunc("/planet/{key}", GetPlanet).Methods("GET")
	router.HandleFunc("/planet", CreatePlanet).Methods("POST")
	router.HandleFunc("/planet/{id}", DeletePlanet).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
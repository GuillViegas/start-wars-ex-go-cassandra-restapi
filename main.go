package main

import (
	"github.com/GuillViegas/start-wars-ex-go-cassandra-restapi/app"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/planet", GetUniverse).Methods("GET")
	router.HandleFunc("/planet/{key}", GetPlanet).Methods("GET")
	router.HandleFunc("/planet", CreatePlanet).Methods("POST")
	router.HandleFunc("/planet/{id}", DeletePlanet).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
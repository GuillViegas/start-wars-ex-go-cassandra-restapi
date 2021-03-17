package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/GuillViegas/start-wars-ex-go-cassandra-restapi/app/handler"
	"github.com/GuillViegas/start-wars-ex-go-cassandra-restapi/app/cassandra"
	"github.com/GuillViegas/start-wars-ex-go-cassandra-restapi/config"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(config *config.Config) {
	session := cassandra.InitCluster(config.Cassandra)

	fmt.Println(config.Cassandra.Host)
	fmt.Println(config.Cassandra.Username)
	fmt.Println(config.Cassandra.Password)
	fmt.Println(config.Cassandra.Keyspace)

	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Router.HandleFunc("/planet", handler.GetUniverse).Methods("GET")
	a.Router.HandleFunc("/planet/{key}", handler.GetPlanet).Methods("GET")
	a.Router.HandleFunc("/planet", handler.CreatePlanet).Methods("POST")
	a.Router.HandleFunc("/planet/{id}", handler.DeletePlanet).Methods("DELETE")
}

func (a *App) Run (host string) {
	log.Fatal(http.ListenAndServe)(host, a.Router))
}

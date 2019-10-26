package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"hackaton-facef-api/api/handler"

	"github.com/gorilla/mux"
)

// App struct ...
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//StartServer ...
func (a *App) StartServer() {
	a.Router = mux.NewRouter()
	s := a.Router.PathPrefix("/api/v1").Subrouter()
	//Health para teste de funcionamento da API
	s.HandleFunc("/health", handler.HealthCheck).Methods(http.MethodGet)

	a.Router.Handle("/api/v1/{_:.*}", a.Router)
	port := 8081
	log.Printf("Starting Server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router))
}

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

	s.HandleFunc("/login", handler.Login).Methods(http.MethodPost)

	s.HandleFunc("/user", handler.InsertUser).Methods(http.MethodPost)
	s.HandleFunc("/user", handler.GetUser).Methods(http.MethodGet)

	s.HandleFunc("/questions", handler.GetQuestion).Methods(http.MethodGet)

	s.HandleFunc("/historic/{id:[0-9]+}", handler.GetHistoric).Methods(http.MethodGet)

	s.HandleFunc("/ranking", handler.GetRanking).Methods(http.MethodGet)

	s.HandleFunc("/point/{point:[0-9]+}/user/{id:[0-9]+}", handler.UpdatePoints).Methods(http.MethodPut)

	a.Router.Handle("/api/v1/{_:.*}", a.Router)
	port := 8081
	log.Printf("Starting Server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), a.Router))
}

package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"

	"filmoteka/internal/app/store"
)

type server struct {
	router *mux.Router
	logger *slog.Logger
	store  store.IStore
}

func NewServer(store store.IStore) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: slog.Default(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/films/{id}", s.handleFilmFind()).Methods("GET")
	s.router.HandleFunc("/films", s.handleFilmCreate()).Methods("POST")
	s.router.HandleFunc("/films", s.handleAllFilms()).Methods("GET")
	s.router.HandleFunc("/films/{id}", s.handleFilmDelete()).Methods("DELETE")
	s.router.HandleFunc("/films/{id}", s.handleFilmUpdate()).Methods("PUT")
	s.router.HandleFunc("/actors/{id}", s.handleActorFind()).Methods("GET")
	s.router.HandleFunc("/actors", s.handleActorCreate()).Methods("POST")
	s.router.HandleFunc("/actors", s.handleAllActors()).Methods("GET")
	s.router.HandleFunc("/actors/{id}", s.handleActorDelete()).Methods("DELETE")
	s.router.HandleFunc("/actors/{id}", s.handleActorUpdate()).Methods("PUT")
}

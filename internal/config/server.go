package config

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krishnapramodaradhi/qlik-api/internal/handler"
	"github.com/krishnapramodaradhi/qlik-api/internal/util"
)

type Server struct {
	listenAddr string
	db         *sql.DB
}

func (s *Server) Run() {
	r := mux.NewRouter()
	u := r.PathPrefix("/api/utility").Subrouter()
	u.HandleFunc("/search", util.WithDB(handler.SearchFilters, s.db)).Methods(http.MethodGet)
	u.HandleFunc("/filters", util.WithDB(handler.FetchFilters, s.db)).Methods(http.MethodGet)
	u.Use(mux.CORSMethodMiddleware(u))

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		util.WriteJSON(w, http.StatusNotFound, map[string]any{"type": "error", "result": "Route or HTTP method is not registered"})
	})

	log.Fatal(http.ListenAndServe(s.listenAddr, r))
}

func NewServer(listenAddr string, db *sql.DB) *Server {
	return &Server{listenAddr: listenAddr, db: db}
}

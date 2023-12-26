package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/post", makeHTTPHandlerFunc(s.handlePost))

	router.HandleFunc("/post/{id}", makeHTTPHandlerFunc(s.handleGetPost))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handlePost(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetPost(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreatePost(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeletePost(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetPost(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]

	return WriteJSON(w, http.StatusOK, id)
}

func (s *APIServer) handleCreatePost(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeletePost(w http.ResponseWriter, r *http.Request) error {
	return nil
}

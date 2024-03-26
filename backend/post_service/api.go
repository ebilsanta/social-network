package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string `json:"error"`
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
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/posts", makeHTTPHandlerFunc(s.handlePost))

	router.HandleFunc("/posts/{id}", makeHTTPHandlerFunc(s.handlePostByID))

	router.HandleFunc("/users/{id}/posts", makeHTTPHandlerFunc(s.handleGetPostsByUserID))

	log.Println("Post service running on port:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handlePost(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetPosts(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreatePost(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetPosts(w http.ResponseWriter, r *http.Request) error {
	posts, err := s.store.GetPosts()

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, posts)
}

func (s *APIServer) handleCreatePost(w http.ResponseWriter, r *http.Request) error {
	createPostReq := CreatePostRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createPostReq); err != nil {
		return err
	}

	post := NewPost(createPostReq.Caption, createPostReq.ImageURL, createPostReq.PosterID)
	dbPost, err := s.store.CreatePost(post)

	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, dbPost)
}

func (s *APIServer) handlePostByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetPostByID(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeletePostByID(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *APIServer) handleGetPostByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	post, err := s.store.GetPostByID(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, post)
}

func (s *APIServer) handleDeletePostByID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	if err := s.store.DeletePost(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}

func (s *APIServer) handleGetPostsByUserID(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	posts, err := s.store.GetPostsByUserID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, posts)

}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id %s", idStr)
	}
	return id, nil
}

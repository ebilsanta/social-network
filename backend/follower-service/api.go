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

	router.HandleFunc("/users/{id}/followers", makeHTTPHandlerFunc(s.handleGetFollowers))

	router.HandleFunc("/users/{id}/following", makeHTTPHandlerFunc(s.handleGetFollowing))

	router.HandleFunc("/users/{followerID}/follow/{followedUserID}", makeHTTPHandlerFunc(s.handleFollow))

	router.HandleFunc("/users/{followerID}/unfollow/{followedUserID}", makeHTTPHandlerFunc(s.handleUnfollow))

	log.Println("Follower service running on port:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleGetFollowers(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	followers, err := s.store.GetFollowers(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, followers)
}

func (s *APIServer) handleGetFollowing(w http.ResponseWriter, r *http.Request) error {
	id, err := getID(r)
	if err != nil {
		return err
	}

	following, err := s.store.GetFollowing(id)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, following)
}

func (s *APIServer) handleFollow(w http.ResponseWriter, r *http.Request) error {
	followerID, followedUserID, err := getFollowerAndFollowedID(r)
	if err != nil {
		return err
	}

	if err := s.store.AddFollower(followerID, followedUserID); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, map[string]string{"message": fmt.Sprintf("User %s has followed user %s", followerID, followedUserID)})
}

func (s *APIServer) handleUnfollow(w http.ResponseWriter, r *http.Request) error {
	followerID, followedUserID, err := getFollowerAndFollowedID(r)
	if err != nil {
		return err
	}

	if err := s.store.DeleteFollower(followerID, followedUserID); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, map[string]string{"message": fmt.Sprintf("User %s has unfollowed user %s", followerID, followedUserID)})
}

func getID(r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, fmt.Errorf("invalid id %s", idStr)
	}
	return id, nil
}

func getFollowerAndFollowedID(r *http.Request) (int, int, error) {
	followerIDStr := mux.Vars(r)["followerID"]
	followerID, err := strconv.Atoi(followerIDStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid followerID %s", followerIDStr)
	}
	followedUserIDStr := mux.Vars(r)["followedUserID"]
	followedUserID, err := strconv.Atoi(followedUserIDStr)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid followerID %s", followedUserIDStr)
	}
	return followerID, followedUserID, nil
}

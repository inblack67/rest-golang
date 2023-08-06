package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type GreetRes struct {
	Hello string `json:"hello"`
}

func (s *APIServer) handleGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		w.Write([]byte("Method not supported"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	res := &GreetRes{
		Hello: "worlds",
	}
	json.NewEncoder(w).Encode(res)
}

func (s *APIServer) handleGetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		w.Write([]byte("Method not supported"))
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Valid integer id is required"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	res, err := s.storage.getPost(id)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("id does not exist"))
		return
	}

	json.NewEncoder(w).Encode(res)
}

func (s *APIServer) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		w.Write([]byte("Method not supported"))
		return
	}

	payload := new(CreatePostPayload)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Invalid payload"))
		return
	}

	fmt.Printf("title =  %s and content = %s", payload.Title, payload.Content)

	// add to db
	post := &Post{
		Title:     payload.Title,
		Content:   payload.Content,
		CreatedAt: time.Now(),
	}

	err = s.storage.persistPost(post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("Post created"))
}

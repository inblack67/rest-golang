package main

import (
	"encoding/json"
	"net/http"
)

type GreetRes struct {
	Hello string `json:"hello"`
}

func (s *APIServer) handleGreet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	res := &GreetRes{
		Hello: "worlds",
	}
	json.NewEncoder(w).Encode(res)
}

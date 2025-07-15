package main

import (
	"encoding/json"
	"net/http"
	"log"
)
func respondwithJSON(w http.ResponseWriter, code int , payload interface{}){
	data,err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(500)
		log.Printf("Failed to parse JSON response: %v",payload)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(data)
}
func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499{
		log.Println("Responding with 500 level",msg)
	}
    type errResponse struct {
		Error string `json:"error"`
	}
	respondwithJSON(w, code,errResponse{
		Error:msg,
	})
}
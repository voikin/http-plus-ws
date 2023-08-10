package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type RandomNumberResponse struct {
	RandomNumber int `json:"randomNumber"`
}

func randomNumberHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10) + 1

	response := RandomNumberResponse{
		RandomNumber: randomNumber,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonResponse)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/random", randomNumberHandler)

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("беда")
	}
}

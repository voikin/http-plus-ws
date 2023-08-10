package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	for {
		randomDelay := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(randomDelay)

		message := "ready to update"
		jsonMessage, _ := json.Marshal(map[string]string{"message": message})

		if err := conn.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
			fmt.Println("Error writing WebSocket message:", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/ws", websocketHandler)

	fmt.Println("WebSocket server listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("бедаа")
	}
}

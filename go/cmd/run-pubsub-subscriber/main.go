package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type PushSubBody struct {
	Message      Message `json:"message"`
	Subscription string  `json:"subscription"`
}

type Message struct {
	Data        []byte    `json:"data"`
	MessageID   string    `json:"messageId"`
	PublishTime time.Time `json:"publishTime"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var pushSubReq PushSubBody
	err := json.NewDecoder(r.Body).Decode(&pushSubReq)
	if err != nil {
		log.Printf("Failed to decode body: %v", err)
		http.Error(w, "Failed to decode body: "+err.Error(), http.StatusInternalServerError)
	}
	log.Printf("data: %s, message: %#v", string(pushSubReq.Message.Data), pushSubReq)
	w.Write([]byte("received request\n"))
}

func run() error {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	return fmt.Errorf("Failed to listen and serve: %w", err)
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

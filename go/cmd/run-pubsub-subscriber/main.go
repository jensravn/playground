package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jensravn/playground/go/internal/gob"
	"github.com/jensravn/playground/go/internal/item"
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
	e, err := gob.DecodeJSON[item.Entity](r.Body)
	if err != nil {
		log.Printf("Failed to decode entity: %v", err)
		http.Error(w, "Failed to decode entity: "+err.Error(), http.StatusInternalServerError)
	}
	log.Printf("message entity: %#v", e)
	w.Write([]byte("received request\n"))
}

func run() error {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s\n", port)
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

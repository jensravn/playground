package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/jensravn/playground/go/internal/gob"
	"github.com/jensravn/playground/go/internal/item"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	psClient, err := pubsub.NewClient(ctx, "jensravn-playground")
	if err != nil {
		fmt.Printf("Failed to create client: %v", err)
		http.Error(w, "Failed to create client: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer psClient.Close()
	topic := psClient.Topic("topic-id")
	ii := []item.Entity{
		{No: "1", Type: "A"},
		{No: "2", Type: "B"},
		{No: "3", Type: "C"},
		{No: "1", Type: "A"},
		{No: "2", Type: "B"},
		{No: "3", Type: "C"},
		{No: "1", Type: "A"},
		{No: "2", Type: "B"},
		{No: "3", Type: "C"},
	}
	log.Printf("Publishing %d messages", len(ii))
	for _, i := range ii {
		log.Printf("Publishing message: %#v", i)
		bb, err := gob.Encode(i)
		if err != nil {
			log.Printf("Failed to encode: %v", err)
			http.Error(w, "Failed to encode: "+err.Error(), http.StatusInternalServerError)
			return
		}
		m := pubsub.Message{Data: bb}
		res := topic.Publish(ctx, &m)
		id, err := res.Get(ctx)
		if err != nil {
			log.Printf("Failed to publish: %v", err)
		} else {
			log.Printf("Published message=%v ID=%v", i, id)
		}
	}
	w.Write([]byte("10 messages published"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	log.Printf("Failed to listen and serve: %v", err)
}

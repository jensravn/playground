package main

import (
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	psClient, err := pubsub.NewClient(r.Context(), "project-id")
	if err != nil {
		http.Error(w, "Failed to create client: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer psClient.Close()
	topic := psClient.Topic("topic-id")
	for i := 1; i <= 100; i++ {
		_ = topic.Publish(r.Context(), &pubsub.Message{Data: []byte(fmt.Sprintf("Number %d", i))})
	}
	w.Write([]byte("100 messages published"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}

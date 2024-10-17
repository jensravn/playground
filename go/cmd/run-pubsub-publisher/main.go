package main

import (
	"fmt"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
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
	for i := 1; i <= 10; i++ {
		res := topic.Publish(ctx, &pubsub.Message{Data: []byte(fmt.Sprintf("Number %d", i))})
		id, err := res.Get(ctx)
		if err != nil {
			fmt.Printf("Failed to publish: %v", err)
		} else {
			fmt.Printf("Published message %d; msg ID: %v\n", i, id)
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
	err := http.ListenAndServe(":"+port, nil)
	fmt.Printf("Failed to listen and serve: %v", err)
}

package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/jensravn/playground/go/internal/shared"
)

// send and receive binary messages with no schema

func main() {
	// send binary data
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "goplayground-pubsubbq")
	if err != nil {
		log.Fatalln("new pubsub client:", err)
	}
	topic := client.Topic("person_binary")
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	person := shared.Person{
		ID:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*shared.PhoneNumber{{
			Number: "555-4321",
			Type:   shared.PhoneTypeHOME,
		}},
	}
	err = enc.Encode(person)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	bb := network.Bytes()
	msg := &pubsub.Message{Data: bb}
	res := topic.Publish(ctx, msg)
	topic.Flush()
	log.Printf("res: %#v", res)

	// receive binary data
	sub := client.Subscription("person_binary-sub")
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		dec := gob.NewDecoder(bytes.NewBuffer(msg.Data))
		var person shared.Person
		err = dec.Decode(&person)
		if err != nil {
			log.Fatal("decode error:", err)
		}
		log.Printf("decoded: %#v", person)
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})
	if err != nil {
		log.Fatal("sub.Receive:", err)
	}
	log.Printf("Received %d messages\n", received)
}

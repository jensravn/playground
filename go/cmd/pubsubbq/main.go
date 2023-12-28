package main

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/jensravn/playground/go/internal/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "goplayground-pubsubbq")
	if err != nil {
		log.Fatalln("new pubsub client:", err)
	}
	topic := client.Topic("person_with_schema")

	person := &pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
		LastUpdated: timestamppb.Now(),
	}
	b, err := proto.Marshal(person)
	// person := Person{
	// 	Id:    1234,
	// 	Name:  "John Doe",
	// 	Email: "jdoe@example.com",
	// 	Phones: []*PhoneNumber{{
	// 		Number: "555-4321",
	// 		Type:   Home,
	// 	}},
	// }
	// b, err := json.Marshal(&person)
	if err != nil {
		log.Fatalln("Failed to create person:", err)
	}
	msg := &pubsub.Message{Data: b}
	res := topic.Publish(ctx, msg)
	topic.Flush()
	log.Printf("res: %#v", res)
}

type Person struct {
	Id     int
	Name   string
	Email  string
	Phones []*PhoneNumber
}
type PhoneNumber struct {
	Number string
	Type   PhoneType
}

type PhoneType int

const (
	Mobile PhoneType = 1
	Home   PhoneType = 2
	Work   PhoneType = 3
)

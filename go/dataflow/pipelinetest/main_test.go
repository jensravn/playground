package main

import (
	"testing"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/passert"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/ptest"
	"github.com/jensravn/playground/go/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestPipeline(t *testing.T) {
	// setup
	beam.Init()
	beamPipeline, scope := beam.NewPipelineWithRoot()

	// given
	persons := []*pb.Person{
		{
			Id:    1234,
			Name:  "John Doe",
			Email: "jdoe@example.com",
			Phones: []*pb.Person_PhoneNumber{
				{Number: "555-4321", Type: pb.Person_HOME},
			},
			LastUpdated: timestamppb.Now(),
		},
	}
	readInput := readInputFn(persons)

	// when
	elements := pipeline(scope, readInput)

	// then
	var anyPersons []interface{} = make([]interface{}, len(persons))
	for i, d := range persons {
		anyPersons[i] = d
	}
	passert.Equals(scope, elements, anyPersons...)
	ptest.RunAndValidate(t, beamPipeline)
}

func readInputFn(persons []*pb.Person) func(beam.Scope) beam.PCollection {
	return func(scope beam.Scope) beam.PCollection {
		var bb [][]byte
		for _, p := range persons {
			b, err := proto.Marshal(p)
			if err != nil {
				panic("failed to marshal person: " + err.Error())
			}
			bb = append(bb, b)
		}
		return beam.CreateList(scope, bb)
	}
}

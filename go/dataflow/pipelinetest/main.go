package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/pubsubio"
	beamLog "github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"github.com/jensravn/playground/go/pb"
	"google.golang.org/protobuf/proto"
)

func init() {
	register.Function3x0(logAndEmit)
	register.Emitter1[*pb.Person]()
}

func main() {
	beam.Init()
	beamPipeline, scope := beam.NewPipelineWithRoot()
	readInput := func(scope beam.Scope) beam.PCollection {
		return pubsubio.Read(scope, "goplayground-pubsubbq", "person_with_schema", nil)
	}

	pipeline(scope, readInput)

	ctx := context.Background()
	err := beamx.Run(ctx, beamPipeline)
	if err != nil {
		log.Fatalf("Failed to run beam pipeline: %v", err)
	}
}

func pipeline(scope beam.Scope, readInput func(scope beam.Scope) beam.PCollection) beam.PCollection {
	messages := readInput(scope)
	stringMsgs := beam.ParDo(scope, protoToString, messages)
	return beam.ParDo(scope, logAndEmit, stringMsgs)
}

func protoToString(message []byte) *pb.Person {
	var person pb.Person
	err := proto.Unmarshal(message, &person)
	if err != nil {
		beamLog.Errorln(context.Background(), fmt.Errorf("failed to unmarshal message: %v", err))
		return nil
	}
	return &person
}

func logAndEmit(ctx context.Context, element *pb.Person, emit func(*pb.Person)) {
	beamLog.Infoln(ctx, "test1")
	beamLog.Infoln(ctx, element.String())
	beamLog.Infoln(ctx, "test2")
	emit(element)
}

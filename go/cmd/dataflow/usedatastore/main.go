package main

import (
	"context"
	"flag"
	"log"

	"cloud.google.com/go/datastore"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	beamLog "github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"github.com/jensravn/playground/go/internal/shared"
)

func init() {
	register.Emitter1[*shared.Person]()
	register.Function3x0(PutInDatastore)
}

func main() {
	flag.Parse()
	beam.Init()
	ctx := context.Background()
	pipeline, scope := beam.NewPipelineWithRoot()

	elements := beam.Create(scope, &shared.Person{
		Name:  "John5",
		Email: "john@john.com",
		Phones: []*shared.PhoneNumber{{
			Number: "12345678",
			Type:   shared.PhoneTypeHOME,
		}},
	})
	beam.ParDo(scope, PutInDatastore, elements)

	err := beamx.Run(ctx, pipeline)
	if err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}

func PutInDatastore(ctx context.Context, element *shared.Person, emit func(*shared.Person)) {
	dsClient, err := datastore.NewClient(ctx, "goplayground-pubsubbq")
	if err != nil {
		beamLog.Errorln(ctx, err)
	}
	defer dsClient.Close()
	k := datastore.NameKey("Person", "John5", nil)
	_, err = dsClient.Put(ctx, k, element)
	if err != nil {
		beamLog.Errorln(ctx, err)
	}
	beamLog.Infoln(ctx, element)
	emit(element)
}

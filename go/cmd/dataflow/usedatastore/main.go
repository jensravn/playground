package main

import (
	"context"
	"flag"

	"cloud.google.com/go/datastore"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"github.com/jensravn/playground/go/internal/shared"
)

var projectID = flag.String("project_id", "", "Google Cloud Project ID")

func init() {
	register.DoFn3x0[context.Context, *shared.Person, func(*shared.Person)](&datastorePutFn{})
	register.Emitter1[*shared.Person]()
}

func main() {
	flag.Parse()
	beam.Init()
	ctx := context.Background()
	if *projectID == "" {
		log.Fatalf(ctx, "flag --project_id must be set")
	}
	pipeline, scope := beam.NewPipelineWithRoot()

	elements := beam.Create(scope, &shared.Person{
		Name:  "John9",
		Email: "john@john.com",
		Phones: []*shared.PhoneNumber{{
			Number: "12345678",
			Type:   shared.PhoneTypeHOME,
		}},
	})
	beam.ParDo(scope, &datastorePutFn{ProjectID: *projectID}, elements)

	err := beamx.Run(ctx, pipeline)
	if err != nil {
		log.Fatalf(ctx, "Failed to execute job: %v", err)
	}
}

type datastorePutFn struct {
	ProjectID string `json:"projectID"`
}

func (f *datastorePutFn) ProcessElement(ctx context.Context, person *shared.Person, emit func(*shared.Person)) {
	dsClient, err := datastore.NewClient(ctx, f.ProjectID)
	if err != nil {
		log.Fatalf(ctx, "Failed to create new datastore client: %v", err)
	}
	defer dsClient.Close()
	k := datastore.NameKey("Person", person.Name, nil)
	_, err = dsClient.Put(ctx, k, person)
	if err != nil {
		log.Errorf(ctx, "Failed to put person to datastore: %v", err)
	}
	emit(person)
}

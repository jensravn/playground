package main

import (
	"context"
	"flag"
	"time"

	"cloud.google.com/go/pubsub/apiv1/pubsubpb"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/graph/window"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/pubsubio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"github.com/jensravn/playground/go/internal/gob"
	"github.com/jensravn/playground/go/internal/item"
)

func init() {
	// register.DoFn1x2[]()
}

func main() {
	flag.Parse()
	beam.Init()
	ctx := context.Background()
	p := beam.NewPipeline()
	s := p.Root()
	data := pubsubio.Read(s, "jensravn-playground", "topic-id", &pubsubio.ReadOptions{WithAttributes: true})
	kvData := beam.ParDo(s, func(message *pubsubpb.PubsubMessage) (string, []byte) {
		e, err := gob.Decode[item.Entity](message.Data)
		if err != nil {
			log.Fatalf(ctx, "Failed to decode entity: %v", err)
		}
		key := e.No + e.Type
		return key, message.Data
	}, data)
	windowed := beam.WindowInto(s, window.NewFixedWindows(30*time.Second), kvData)
	combined := beam.CombinePerKey(s, func(v, _ []byte) []byte { return v }, windowed)
	deduplicated := beam.DropKey(s, combined)
	pubsubio.Write(s, "jensravn-playground", "topic-id-2", deduplicated)
	if err := beamx.Run(ctx, p); err != nil {
		log.Fatalf(ctx, "Failed to execute job: %v", err)
	}
}

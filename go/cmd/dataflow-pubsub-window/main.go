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
)

func init() {}

func main() {
	flag.Parse()
	beam.Init()
	ctx := context.Background()
	p := beam.NewPipeline()
	s := p.Root()

	messages := pubsubio.Read(s, "jensravn-playground", "topic-id", &pubsubio.ReadOptions{WithAttributes: true})
	kvPairs := beam.ParDo(s, func(m *pubsubpb.PubsubMessage) (string, *pubsubpb.PubsubMessage) {
		return "test", m
	}, messages)
	windowed := beam.WindowInto(s, window.NewFixedWindows(30*time.Second), kvPairs)
	combined := beam.CombinePerKey(s, func(m, _ *pubsubpb.PubsubMessage) *pubsubpb.PubsubMessage {
		return m
	}, windowed)
	deduplicated := beam.DropKey(s, combined)
	pubsubio.Write(s, "jensravn-playground", "topic-id-2", deduplicated)

	if err := beamx.Run(ctx, p); err != nil {
		log.Fatalf(ctx, "Failed to execute job: %v", err)
	}
}

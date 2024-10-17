package main

import (
	"context"
	"flag"
	"time"

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
	data := pubsubio.Read(s, "jensravn-playground", "topic-id", nil)
	kvData := beam.ParDo(s, func(elm []byte) ([]byte, []byte) { return []byte("test"), elm }, data)
	windowed := beam.WindowInto(s, window.NewFixedWindows(30*time.Second), kvData)
	combined := beam.CombinePerKey(s, func(v, _ []byte) []byte { return v }, windowed)
	deduplicated := beam.DropKey(s, combined)
	pubsubio.Write(s, "jensravn-playground", "topic-id-2", deduplicated)
	if err := beamx.Run(ctx, p); err != nil {
		log.Fatalf(ctx, "Failed to execute job: %v", err)
	}
}

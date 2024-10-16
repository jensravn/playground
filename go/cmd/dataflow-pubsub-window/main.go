package main

import (
	"context"
	"flag"
	"time"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/graph/window"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/pubsubio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
)

var (
	expansionAddr = flag.String("expansion_addr", "localhost:1234",
		"Address of Expansion Service. If not specified, attempts to automatically start an appropriate expansion service.")
	bootstrapServers = flag.String("bootstrap_servers", "kafka_server:9092",
		"(Required) URL of the bootstrap servers for the Kafka cluster. Should be accessible by the runner.")
	topic = flag.String("topic", "kafka_taxirides_realtime", "Kafka topic to write to and read from.")
)

func init() {
	register.DoFn2x0[context.Context, []byte](&LogFn{})
}

type LogFn struct{}

func (fn *LogFn) ProcessElement(ctx context.Context, elm []byte) {
	log.Infof(ctx, "Ride info: %v", string(elm))
}

func (fn *LogFn) FinishBundle() {
	time.Sleep(2 * time.Second)
}

func main() {
	flag.Parse()
	beam.Init()

	ctx := context.Background()

	p := beam.NewPipeline()
	s := p.Root()

	data := pubsubio.Read(s, "pubsub-public-data", "taxirides-realtime", nil)
	kvData := beam.ParDo(s, func(elm []byte) ([]byte, []byte) { return []byte(""), elm }, data)
	windowed := beam.WindowInto(s, window.NewFixedWindows(60*time.Second), kvData)
	pubsubio.Write(s, "pubsub-public-data", "taxirides-realtime", windowed)

	if err := beamx.Run(ctx, p); err != nil {
		log.Fatalf(ctx, "Failed to execute job: %v", err)
	}
}

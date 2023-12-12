// From https://github.com/apache/beam/blob/master/sdks/go/examples/wordcount/wordcount.go
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/transforms/stats"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
)

// Using Parameterizable PipelineOptions
var (
	input  = flag.String("input", "gs://apache-beam-samples/shakespeare/kinglear.txt", "File(s) to read.")
	output = flag.String("output", "", "Output file (required).")
)

func init() {
	register.DoFn3x0[context.Context, string, func(string)](&extractFn{})
	register.Emitter1[string]()
}

var (
	wordRE          = regexp.MustCompile(`[a-zA-Z]+('[a-z])?`)
	empty           = beam.NewCounter("extract", "emptyLines")
	smallWordLength = flag.Int("small_word_length", 9, "length of small words (default: 9)")
	smallWords      = beam.NewCounter("extract", "smallWords")
	lineLen         = beam.NewDistribution("extract", "lineLenDistro")
)

type extractFn struct {
	SmallWordLength int `json:"smallWordLength"`
}

func (f *extractFn) ProcessElement(ctx context.Context, line string, emit func(string)) {
	lineLen.Update(ctx, int64(len(line)))
	if len(strings.TrimSpace(line)) == 0 {
		empty.Inc(ctx, 1)
	}
	for _, word := range wordRE.FindAllString(line, -1) {
		if len(word) < f.SmallWordLength {
			smallWords.Inc(ctx, 1)
		}
		emit(word)
	}
}

func formatFn(w string, c int) string {
	return fmt.Sprintf("%s: %v", w, c)
}

// Creating Composite Transforms
func CountWords(s beam.Scope, lines beam.PCollection) beam.PCollection {
	s = s.Scope("CountWords")
	col := beam.ParDo(s, &extractFn{SmallWordLength: *smallWordLength}, lines)
	return stats.Count(s, col)
}

func main() {
	flag.Parse()
	beam.Init()
	if *output == "" {
		log.Fatal("No output provided")
	}
	p := beam.NewPipeline()
	s := p.Root()

	lines := textio.Read(s, *input)
	counted := CountWords(s, lines)

	// Applying ParDo with an explicit DoFn
	formatted := beam.ParDo(s, formatFn, counted)
	textio.Write(s, *output, formatted)
	if err := beamx.Run(context.Background(), p); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}

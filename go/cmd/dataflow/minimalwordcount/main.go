package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/gcs"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/direct"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/transforms/stats"
)

// MinimalWordCount demonstrates the basic principles involved in building a pipeline.
// https://beam.apache.org/get-started/wordcount-example/#minimalwordcount-example

var wordRE = regexp.MustCompile(`[a-zA-Z]+('[a-z])?`)

func main() {
	// Creating the Pipeline
	beam.Init()
	p := beam.NewPipeline()
	s := p.Root()

	// Transform - Input to PCollection
	lines := textio.Read(s, "gs://apache-beam-samples/shakespeare/kinglear.txt")

	// Transform - ParDo (scope, DoFn, PCollection) PCollection
	words := beam.ParDo(s, func(line string, emit func(string)) {
		for _, word := range wordRE.FindAllString(line, -1) {
			emit(word)
		}
	}, lines)

	// Transform - SDK-provided transform (key/value pairs)
	counted := stats.Count(s, words)

	// Transform - key/value pairs to printable string
	formatted := beam.ParDo(s, func(w string, c int) string {
		return fmt.Sprintf("%s: %v", w, c)
	}, counted)

	// Transform PCollection to output
	textio.Write(s, "wordcounts.txt", formatted)

	// Running the Pipeline with direct runner
	direct.Execute(context.Background(), p)
}

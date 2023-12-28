// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// https://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or https://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/pubsubio"
	beamLog "github.com/apache/beam/sdks/v2/go/pkg/beam/log"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"github.com/jensravn/playground/go/internal/pb"
	"google.golang.org/protobuf/proto"
)

var input_text = flag.String("input-text", "default input text", "Input text to print.")

func init() {
	// DoFns should be registered with Beam to be available in distributed runners.
	register.Function3x0(logAndEmit)
	register.Emitter1[string]()
}

// You can also access the Context and "emit" zero or more values like FlatMap.
func logAndEmit(ctx context.Context, element string, emit func(string)) {
	beamLog.Infoln(ctx, "test1")
	beamLog.Infoln(ctx, element)
	beamLog.Infoln(ctx, "test2")
	emit(element)
}

func myPipeline(scope beam.Scope, input_text string) beam.PCollection {
	messages := pubsubio.Read(scope, "goplayground-pubsubbq", "person_with_schema", nil)

	stringMsgs := beam.ParDo(scope, func(message []byte) string {
		var person pb.Person
		err := proto.Unmarshal(message, &person)
		if err != nil {
			beamLog.Errorln(context.Background(), fmt.Errorf("failed to unmarshal message: %v", err))
			return ""
		}
		return person.String()
	}, messages)

	return beam.ParDo(scope, logAndEmit, stringMsgs)
}

func main() {
	flag.Parse()
	beam.Init()

	ctx := context.Background()
	pipeline, scope := beam.NewPipelineWithRoot()
	myPipeline(scope, *input_text)

	// Run the pipeline. You can specify your runner with the --runner flag.
	err := beamx.Run(ctx, pipeline)
	if err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}

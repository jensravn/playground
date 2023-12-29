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

// Read messages published to a Pub/Sub topic
// https://cloud.google.com/pubsub/docs/stream-messages-dataflow
// https://github.com/apache/beam/blob/master/sdks/go/pkg/beam/io/pubsubio/pubsubio.go
// https://pkg.go.dev/github.com/apache/beam/sdks/go/pkg/beam/io/pubsubio

func init() {
	register.Function3x0(logAndEmit)
	register.Emitter1[string]()
}

func main() {
	beam.Init()
	ctx := context.Background()
	pipeline, scope := beam.NewPipelineWithRoot()
	myPipeline(scope)
	err := beamx.Run(ctx, pipeline)
	if err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}

func myPipeline(scope beam.Scope) beam.PCollection {
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

func logAndEmit(ctx context.Context, element string, emit func(string)) {
	beamLog.Infoln(ctx, "test1")
	beamLog.Infoln(ctx, element)
	beamLog.Infoln(ctx, "test2")
	emit(element)
}

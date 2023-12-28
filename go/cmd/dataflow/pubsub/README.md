# Dataflow: Pub/Sub

https://cloud.google.com/pubsub/docs/stream-messages-dataflow

## Read messages published to a Pub/Sub topic

- https://github.com/apache/beam/blob/master/sdks/go/pkg/beam/io/pubsubio/pubsubio.go
- https://pkg.go.dev/github.com/apache/beam/sdks/go/pkg/beam/io/pubsubio

```bash
go run . \
    --runner dataflow \
    --project PROJECT_ID \
    --region DATAFLOW_REGION \
    --staging_location gs://STORAGE_BUCKET/binaries/
```

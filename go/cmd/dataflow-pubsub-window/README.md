# Dataflow Pub/Sub Window

This Dataflow gets messages from run-pubsub-publisher and sends windowed messages to run-pubsub-subscriber.

The 3 services need `topic-id`, `topic-id-2` and `subscription-id-2`.

```bash
# Run in Dataflow
go run . \
    --runner dataflow \
    --project jensravn-playground \
    --region europe-west1 \
    --temp_location gs://jensravn-playground-dataflow/tmp/ \
    --staging_location gs://jensravn-playground-dataflow/binaries/
```
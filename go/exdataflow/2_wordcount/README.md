# Wordcount example 2 of 5 - WordCount

https://beam.apache.org/get-started/wordcount-example/#wordcount-example

https://cloud.google.com/dataflow/docs/quickstarts/create-pipeline-go

## Log

1. Copy wordcount example from apache/beam/sdks/go/
2. Run `go mod tidy` to add `github.com/apache/beam/sdks/v2` to `go.mod`
3. Direct runner

```
go run . --output output.txt`
```

4. Dataflow runner

```
go run . \
    --input gs://dataflow-samples/shakespeare/kinglear.txt \
    --output gs://<your-gcs-bucket>/counts \
    --runner dataflow \
    --project your-gcp-project \
    --region your-gcp-region \
    --temp_location gs://<your-gcs-bucket>/tmp/ \
    --staging_location gs://<your-gcs-bucket>/binaries/ \
    --worker_harness_container_image=apache/beam_go_sdk:latest
```

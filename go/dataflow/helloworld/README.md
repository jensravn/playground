# Dataflow: Hello World!

https://github.com/apache/beam-starter-go

```bash
# Run locally using direct runner
go run .

# Run on Dataflow
go run . \
    --runner dataflow \
    --project PROJECT_ID \
    --region DATAFLOW_REGION \
    --staging_location gs://STORAGE_BUCKET/binaries/
```

# Dataflow

```bash
# Run on Prism (local portable runner)
go run .

# Run on Dataflow
go run . \
    --runner dataflow \
    --project PROJECT_ID \
    --region DATAFLOW_REGION \
    --staging_location gs://STORAGE_BUCKET/binaries/
```

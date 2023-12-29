# Dataflow

## Run pipeline

```bash
# Run on Prism (local portable runner)
go run .

# Run on Dataflow
go run . \
    --runner dataflow \
    --project $PROJECT_ID \
    --region $DATAFLOW_REGION \
    --staging_location gs://$STORAGE_BUCKET/binaries/
```

## Troubleshooting

### Flags

**ERROR:** command line flags such as `go run . --runner dataflow` are not being picked up.

**FIX:** add `flags.Parse()` to go code.

### Zone/Region

**ERROR:** Startup of the worker pool in zone europe-west3-a failed to bring up any of the desired 1 workers. Please refer to https://cloud.google.com/dataflow/docs/guides/common-errors#worker-pool-failure for help troubleshooting. ZONE_RESOURCE_POOL_EXHAUSTED: Instance 'go-job-1-...-...-...-harness-...' creation failed: The zone 'projects/.../zones/europe-west3-a' does not have enough resources available to fulfill the request. Try a different zone, or try again later.

**FIX:** change to run in `--region europe-west1`

# Dataflow

## Run pipeline

```bash
# Run on Prism (local portable runner)
go run . --project_id $PROJECT_ID

# Run on Dataflow
go run . \
    --project_id $PROJECT_ID \
    --runner dataflow \
    --project $PROJECT_ID \
    --region $DATAFLOW_REGION \
    --staging_location gs://$STORAGE_BUCKET/binaries/
```

## Links

- [Apache Beam: Videos and Podcasts](https://beam.apache.org/get-started/resources/videos-and-podcasts/)
- [Apache Beam: Programming Guide](https://beam.apache.org/documentation/programming-guide/)
- [Skills Boost - Dataflow: Foundations](https://www.cloudskillsboost.google/course_templates/218)

## Troubleshooting

### Flags

**ERROR:** command line flags such as `go run . --runner dataflow` are not being picked up.

**FIX:** add `flags.Parse()` to go code.

### Zone/Region

**ERROR:** Startup of the worker pool in zone europe-west3-a failed to bring up any of the desired 1 workers. Please refer to https://cloud.google.com/dataflow/docs/guides/common-errors#worker-pool-failure for help troubleshooting. ZONE_RESOURCE_POOL_EXHAUSTED: Instance 'go-job-1-...-...-...-harness-...' creation failed: The zone 'projects/.../zones/europe-west3-a' does not have enough resources available to fulfill the request. Try a different zone, or try again later.

**FIX:** change to run in `--region europe-west1`

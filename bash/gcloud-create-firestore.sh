#!/bin/bash

gcloud alpha firestore databases create \
    --database=my-firestore \
    --location=eur3 \
    --type=firestore-native

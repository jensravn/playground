#!/bin/bash

while getopts s: flag; do
    case "${flag}" in
    s) service_account_email=${OPTARG} ;;
    esac
done

echo "service_account_email: $service_account_email"

gcloud auth application-default login --impersonate-service-account $service_account_email

# ./gcloud-local-dev.sh -s ...@...iam.gserviceaccount.com

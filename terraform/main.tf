terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "6.36.1"
    }
  }
}

locals {
  project_id     = "jensravn-tfplayground"
  project_number = "776521789914"
  person_schema  = "person_schema"
}

provider "google" {
  project = local.project_id
  region  = "europe-west1"
  zone    = "europe-west1-b"
}

resource "google_pubsub_schema" "person_schema" {
  name       = local.person_schema
  type       = "PROTOCOL_BUFFER"
  definition = file("./person.proto")
}

resource "google_pubsub_topic" "person_topic" {
  name       = "person"
  depends_on = [google_pubsub_schema.person_schema]
  schema_settings {
    schema   = "projects/${local.project_id}/schemas/${local.person_schema}"
    encoding = "BINARY"
  }
}

resource "google_bigquery_dataset" "pubsub_msgs" {
  dataset_id = "pubsub_msgs"
  location   = "europe-west1"
}

resource "google_bigquery_table" "person_msgs" {
  deletion_protection = false
  table_id            = "person_msgs"
  dataset_id          = google_bigquery_dataset.pubsub_msgs.dataset_id
  schema              = file("./person.bqschema.json")
}

resource "google_project_iam_member" "viewer" {
  project = local.project_id
  role    = "roles/bigquery.metadataViewer"
  member  = "serviceAccount:service-${local.project_number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

resource "google_project_iam_member" "editor" {
  project = local.project_id
  role    = "roles/bigquery.dataEditor"
  member  = "serviceAccount:service-${local.project_number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

resource "google_pubsub_subscription" "person_sub_bq" {
  name  = "person_bigquery_subscription"
  topic = google_pubsub_topic.person_topic.name
  bigquery_config {
    table            = "${google_bigquery_table.person_msgs.project}.${google_bigquery_table.person_msgs.dataset_id}.${google_bigquery_table.person_msgs.table_id}"
    use_topic_schema = true
    write_metadata   = true
  }
  depends_on = [google_project_iam_member.viewer, google_project_iam_member.editor]
}

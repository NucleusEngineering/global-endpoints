locals {
  project        = "stamer-experiments"
  organization   = "780482598012"
  prefix = "global-endpoint"
}

terraform {
  backend "gcs" {
    bucket = "stamer-experiments-global-endpoints-tf-state"
    prefix = "terraform/state"
  }
}

provider "google" {
  project = local.project
}

provider "google-beta" {
  project = local.project
}

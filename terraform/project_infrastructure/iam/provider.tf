terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }
  backend "gcs" {
    bucket  = "wander-project-remote-tfstate-bucket"
    prefix  = "project/iam"
  }
}
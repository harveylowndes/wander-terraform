terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
  }
  backend "gcs" {
    bucket  = "wander-project-remote-tfstate-bucket"
    prefix  = "application"
  }
}

# This makes ops a little annoying
data "google_client_config" "current" {
}

data "google_container_cluster" "primary" {
  name     = "primary-cluster"
  location = terraform.workspace
  project  = var.project_id
}

provider "helm" {
  kubernetes {
    token = data.google_client_config.current.access_token
    host  = data.google_container_cluster.primary.endpoint
    cluster_ca_certificate = base64decode(data.google_container_cluster.primary.master_auth[0].cluster_ca_certificate)
  }
}
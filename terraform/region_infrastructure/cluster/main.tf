locals {
    region  = terraform.workspace
}

data "google_service_account" "cluster_service_account" {
  account_id = "primary-cluster-user"
  project  = var.project_id
}

resource "google_container_cluster" "primary" {
  name     = "primary-cluster"
  location = local.region
  project  = var.project_id

  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "primary" {
  name       = "primary-node-pool"
  location = local.region
  project  = var.project_id

  cluster    = google_container_cluster.primary.name
  node_count = 1

  node_config {
    preemptible  = true
    machine_type = "e2-medium"

    service_account = data.google_service_account.cluster_service_account.email
    oauth_scopes    = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]
  }
}
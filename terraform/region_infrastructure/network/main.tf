locals {
    region  = terraform.workspace
}

data "google_compute_network" "vpc" {
    name    = "${var.environment}-primary-vpc"
    project = var.project_id
}

resource "google_compute_subnetwork" "primary_cluster" {
    name          = "primary-cluster-subnet"
    project       = var.project_id

    ip_cidr_range = "10.0.0.0/22"
    region        = local.region

    network       = data.google_compute_network.vpc.id
}
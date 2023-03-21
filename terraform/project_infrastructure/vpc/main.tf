resource "google_compute_network" "vpc_network" {
  project                 = var.project_id
  name                    = "${var.environment}-primary-vpc"
  description             = "Primary GKE cluster VPC."
  auto_create_subnetworks = false
  delete_default_routes_on_create = false
}
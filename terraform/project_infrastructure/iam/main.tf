resource "google_service_account" "cluster_service_account" {
  account_id   = "primary-cluster-user"
  display_name = "Primary Cluster Service Account"
  project      = var.project_id
}

resource "google_project_iam_member" "allow_image_pull" {
  project = var.project_id
  role   = "roles/artifactregistry.reader"
  member = "serviceAccount:${google_service_account.cluster_service_account.email}"
}

# Outputs are for testing purposes
output "cluster_service_account_email" {
  value = google_service_account.cluster_service_account.email
}

output "cluster_service_account_id" {
  value = google_service_account.cluster_service_account.id
}
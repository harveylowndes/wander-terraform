resource "google_storage_bucket" "tfstate_bucket" {
  name          = "wander-project-remote-tfstate-bucket"
  project       = var.project_id
  force_destroy = true # TODO (hlowndes) Set back to false when finished.
  location      = "US"
  storage_class = "STANDARD"
}

# TODO (hlowndes) Check object lifestyle states

output "bucket_url" {
  value = google_storage_bucket.tfstate_bucket.url
}
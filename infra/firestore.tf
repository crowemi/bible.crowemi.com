resource "google_firestore_database" "crowemi-bible" {
  project     = var.gcp_project_id
  name        = local.service
  location_id = var.gcp_region
  type        = "FIRESTORE_NATIVE"
}

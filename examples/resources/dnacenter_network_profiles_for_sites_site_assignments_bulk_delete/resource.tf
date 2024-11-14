
resource "dnacenter_network_profiles_for_sites_site_assignments_bulk_delete" "example" {
  provider   = dnacenter
  profile_id = "string"
  site_id    = "string"
  parameters {

  }
}

output "dnacenter_network_profiles_for_sites_site_assignments_bulk_delete_example" {
  value = dnacenter_network_profiles_for_sites_site_assignments_bulk_delete.example
}
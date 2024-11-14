
resource "dnacenter_network_profiles_for_sites_site_assignments_bulk_create" "example" {
  provider   = dnacenter
  profile_id = "string"
  parameters {

    items = []
  }
}

output "dnacenter_network_profiles_for_sites_site_assignments_bulk_create_example" {
  value = dnacenter_network_profiles_for_sites_site_assignments_bulk_create.example
}

resource "dnacenter_network_profiles_for_sites_site_assignments" "example" {
  provider = dnacenter

  parameters {

    id         = "string"
    profile_id = "string"
  }
}

output "dnacenter_network_profiles_for_sites_site_assignments_example" {
  value = dnacenter_network_profiles_for_sites_site_assignments.example
}

data "dnacenter_network_profiles_for_sites_site_assignments" "example" {
  provider   = dnacenter
  limit      = 1
  offset     = 1
  profile_id = "string"
}

output "dnacenter_network_profiles_for_sites_site_assignments_example" {
  value = data.dnacenter_network_profiles_for_sites_site_assignments.example.items
}

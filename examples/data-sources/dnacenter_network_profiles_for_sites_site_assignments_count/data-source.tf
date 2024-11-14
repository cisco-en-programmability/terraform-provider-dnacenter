
data "dnacenter_network_profiles_for_sites_site_assignments_count" "example" {
  provider   = dnacenter
  profile_id = "string"
}

output "dnacenter_network_profiles_for_sites_site_assignments_count_example" {
  value = data.dnacenter_network_profiles_for_sites_site_assignments_count.example.item
}

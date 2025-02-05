
data "dnacenter_network_profiles_for_sites_profile_id_templates_count" "example" {
  provider   = dnacenter
  profile_id = "string"
}

output "dnacenter_network_profiles_for_sites_profile_id_templates_count_example" {
  value = data.dnacenter_network_profiles_for_sites_profile_id_templates_count.example.item
}

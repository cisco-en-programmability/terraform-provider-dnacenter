
data "dnacenter_network_profiles_for_sites_profile_id_templates" "example" {
  provider   = dnacenter
  profile_id = "string"
}

output "dnacenter_network_profiles_for_sites_profile_id_templates_example" {
  value = data.dnacenter_network_profiles_for_sites_profile_id_templates.example.items
}

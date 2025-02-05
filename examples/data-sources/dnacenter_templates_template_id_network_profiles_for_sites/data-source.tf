
data "dnacenter_templates_template_id_network_profiles_for_sites" "example" {
  provider    = dnacenter
  template_id = "string"
}

output "dnacenter_templates_template_id_network_profiles_for_sites_example" {
  value = data.dnacenter_templates_template_id_network_profiles_for_sites.example.item
}

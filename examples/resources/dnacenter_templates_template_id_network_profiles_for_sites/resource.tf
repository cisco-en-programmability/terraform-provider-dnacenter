
resource "dnacenter_templates_template_id_network_profiles_for_sites" "example" {
  provider = dnacenter

  parameters {

    profile_id  = "string"
    template_id = "string"
  }
}

output "dnacenter_templates_template_id_network_profiles_for_sites_example" {
  value = dnacenter_templates_template_id_network_profiles_for_sites.example
}

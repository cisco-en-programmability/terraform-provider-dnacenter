
resource "dnacenter_disassociate_site_to_network_profile" "example" {
  provider = dnacenter
  parameters {

    network_profile_id = "string"
    site_id            = "string"
  }
}

output "dnacenter_disassociate_site_to_network_profile_example" {
  value = dnacenter_disassociate_site_to_network_profile.example
}
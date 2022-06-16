
resource "dnacenter_associate_site_to_network_profile" "example" {
  provider = dnacenter
  parameters {

    network_profile_id = "string"
    site_id            = "string"
  }
}

output "dnacenter_associate_site_to_network_profile_example" {
  value = dnacenter_associate_site_to_network_profile.example
}

resource "dnacenter_associate_site_to_network_profile" "example" {
  provider           = dnacenter
  network_profile_id = "string"
  site_id            = "string"
  parameters {

  }
}

output "dnacenter_associate_site_to_network_profile_example" {
  value = dnacenter_associate_site_to_network_profile.example
}
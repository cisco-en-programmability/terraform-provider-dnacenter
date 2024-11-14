
resource "dnacenter_network_devices_network_profiles_for_sites" "example" {
  provider = dnacenter

  parameters {

    id = "string"
  }
}

output "dnacenter_network_devices_network_profiles_for_sites_example" {
  value = dnacenter_network_devices_network_profiles_for_sites.example
}
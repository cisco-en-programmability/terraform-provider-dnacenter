
data "dnacenter_network_devices_network_profiles_for_sites_count" "example" {
  provider = dnacenter
  type     = "string"
}

output "dnacenter_network_devices_network_profiles_for_sites_count_example" {
  value = data.dnacenter_network_devices_network_profiles_for_sites_count.example.item
}

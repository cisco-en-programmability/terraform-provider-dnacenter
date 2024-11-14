
data "dnacenter_network_devices_network_profiles_for_sites" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
  order    = "string"
  sort_by  = "string"
  type     = "string"
}

output "dnacenter_network_devices_network_profiles_for_sites_example" {
  value = data.dnacenter_network_devices_network_profiles_for_sites.example.items
}

data "dnacenter_network_devices_network_profiles_for_sites" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_devices_network_profiles_for_sites_example" {
  value = data.dnacenter_network_devices_network_profiles_for_sites.example.item
}

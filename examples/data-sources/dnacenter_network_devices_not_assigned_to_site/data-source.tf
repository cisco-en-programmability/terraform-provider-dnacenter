
data "dnacenter_network_devices_not_assigned_to_site" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
}

output "dnacenter_network_devices_not_assigned_to_site_example" {
  value = data.dnacenter_network_devices_not_assigned_to_site.example.item
}

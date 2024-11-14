
data "dnacenter_network_devices_not_assigned_to_site_count" "example" {
  provider = dnacenter
}

output "dnacenter_network_devices_not_assigned_to_site_count_example" {
  value = data.dnacenter_network_devices_not_assigned_to_site_count.example.item
}

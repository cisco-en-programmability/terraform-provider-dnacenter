
data "dnacenter_network_devices_assigned_to_site_count" "example" {
  provider = dnacenter
  site_id  = "string"
}

output "dnacenter_network_devices_assigned_to_site_count_example" {
  value = data.dnacenter_network_devices_assigned_to_site_count.example.item
}

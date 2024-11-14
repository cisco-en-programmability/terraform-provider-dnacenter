
data "dnacenter_network_devices_assigned_to_site" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
  site_id  = "string"
}

output "dnacenter_network_devices_assigned_to_site_example" {
  value = data.dnacenter_network_devices_assigned_to_site.example.items
}

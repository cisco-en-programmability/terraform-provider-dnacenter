
data "dnacenter_network_devices_assigned_to_site_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_devices_assigned_to_site_id_example" {
  value = data.dnacenter_network_devices_assigned_to_site_id.example.item
}
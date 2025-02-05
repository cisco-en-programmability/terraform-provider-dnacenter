
data "dnacenter_network_devices_id" "example" {
  provider = dnacenter
  id       = "string"
  views    = "string"
}

output "dnacenter_network_devices_id_example" {
  value = data.dnacenter_network_devices_id.example.item
}

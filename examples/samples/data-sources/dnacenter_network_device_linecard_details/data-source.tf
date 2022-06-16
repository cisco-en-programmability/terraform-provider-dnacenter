
data "dnacenter_network_device_linecard_details" "example" {
  provider    = dnacenter
  device_uuid = "string"
}

output "dnacenter_network_device_linecard_details_example" {
  value = data.dnacenter_network_device_linecard_details.example.items
}

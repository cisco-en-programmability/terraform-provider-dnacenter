
data "dnacenter_network_device_supervisor_card_details" "example" {
  provider    = dnacenter
  device_uuid = "string"
}

output "dnacenter_network_device_supervisor_card_details_example" {
  value = data.dnacenter_network_device_supervisor_card_details.example.items
}

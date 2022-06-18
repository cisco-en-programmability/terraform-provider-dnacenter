
data "dnacenter_network_device_chassis_details" "example" {
  provider  = dnacenter
  device_id = "string"
}

output "dnacenter_network_device_chassis_details_example" {
  value = data.dnacenter_network_device_chassis_details.example.items
}

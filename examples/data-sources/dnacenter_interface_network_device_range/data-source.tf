
data "dnacenter_interface_network_device_range" "example" {
  provider          = dnacenter
  device_id         = "string"
  records_to_return = 1
  start_index       = 1
}

output "dnacenter_interface_network_device_range_example" {
  value = data.dnacenter_interface_network_device_range.example.items
}

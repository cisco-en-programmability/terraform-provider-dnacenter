
data "dnacenter_network_device_range" "example" {
  provider          = dnacenter
  records_to_return = 1
  start_index       = 1
}

output "dnacenter_network_device_range_example" {
  value = data.dnacenter_network_device_range.example.items
}

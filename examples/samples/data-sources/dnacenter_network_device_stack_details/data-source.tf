
data "dnacenter_network_device_stack_details" "example" {
  provider  = dnacenter
  device_id = "string"
}

output "dnacenter_network_device_stack_details_example" {
  value = data.dnacenter_network_device_stack_details.example.item
}

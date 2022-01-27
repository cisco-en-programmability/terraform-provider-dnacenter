
data "dnacenter_interface_network_device" "example" {
  provider  = dnacenter
  device_id = "string"
}

output "dnacenter_interface_network_device_example" {
  value = data.dnacenter_interface_network_device.example.items
}

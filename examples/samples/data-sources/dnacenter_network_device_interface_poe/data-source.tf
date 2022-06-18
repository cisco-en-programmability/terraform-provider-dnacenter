
data "dnacenter_network_device_interface_poe" "example" {
  provider            = dnacenter
  device_uuid         = "string"
  interface_name_list = "string"
}

output "dnacenter_network_device_interface_poe_example" {
  value = data.dnacenter_network_device_interface_poe.example.items
}

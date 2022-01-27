
data "dnacenter_interface_network_device_detail" "example" {
  provider  = dnacenter
  device_id = "string"
  name      = "string"
}

output "dnacenter_interface_network_device_detail_example" {
  value = data.dnacenter_interface_network_device_detail.example.item
}

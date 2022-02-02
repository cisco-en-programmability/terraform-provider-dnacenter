
data "dnacenter_network_device_interface_neighbor" "example" {
  provider       = dnacenter
  device_uuid    = "string"
  interface_uuid = "string"
}

output "dnacenter_network_device_interface_neighbor_example" {
  value = data.dnacenter_network_device_interface_neighbor.example.item
}

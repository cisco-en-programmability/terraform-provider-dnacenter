
data "dnacenter_network_device_vlan" "example" {
  provider       = dnacenter
  id             = "string"
  interface_type = "string"
}

output "dnacenter_network_device_vlan_example" {
  value = data.dnacenter_network_device_vlan.example.items
}

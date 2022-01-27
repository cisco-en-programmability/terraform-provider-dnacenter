
data "dnacenter_device_interface_ospf" "example" {
  provider = dnacenter
}

output "dnacenter_device_interface_ospf_example" {
  value = data.dnacenter_device_interface_ospf.example.items
}

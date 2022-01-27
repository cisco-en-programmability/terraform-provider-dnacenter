
data "dnacenter_device_interface_isis" "example" {
  provider = dnacenter
}

output "dnacenter_device_interface_isis_example" {
  value = data.dnacenter_device_interface_isis.example.items
}

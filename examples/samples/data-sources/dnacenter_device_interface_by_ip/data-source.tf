
data "dnacenter_device_interface_by_ip" "example" {
  provider   = dnacenter
  ip_address = "string"
}

output "dnacenter_device_interface_by_ip_example" {
  value = data.dnacenter_device_interface_by_ip.example.items
}

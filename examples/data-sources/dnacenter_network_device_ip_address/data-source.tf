
data "dnacenter_network_device_ip_address" "example" {
  provider   = dnacenter
  ip_address = "string"
}

output "dnacenter_network_device_ip_address_example" {
  value = data.dnacenter_network_device_ip_address.example.item
}

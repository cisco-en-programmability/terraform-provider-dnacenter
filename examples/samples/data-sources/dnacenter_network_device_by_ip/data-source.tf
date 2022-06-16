
data "dnacenter_network_device_by_ip" "example" {
  provider   = dnacenter
  ip_address = "string"
}

output "dnacenter_network_device_by_ip_example" {
  value = data.dnacenter_network_device_by_ip.example.item
}

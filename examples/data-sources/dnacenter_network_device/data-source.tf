
data "dnacenter_network_device" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_network_device_example" {
  value = data.dnacenter_network_device.example.item
}


data "dnacenter_network_device_poe" "example" {
  provider    = dnacenter
  device_uuid = "string"
}

output "dnacenter_network_device_poe_example" {
  value = data.dnacenter_network_device_poe.example.item
}

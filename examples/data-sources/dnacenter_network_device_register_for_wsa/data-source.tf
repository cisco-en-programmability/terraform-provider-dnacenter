
data "dnacenter_network_device_register_for_wsa" "example" {
  provider      = dnacenter
  macaddress    = "string"
  serial_number = "string"
}

output "dnacenter_network_device_register_for_wsa_example" {
  value = data.dnacenter_network_device_register_for_wsa.example.item
}

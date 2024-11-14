
resource "dnacenter_network_device_management_address_update" "example" {
  provider = dnacenter
  deviceid = "string"
  parameters {

    new_ip = "string"
  }
}

output "dnacenter_network_device_management_address_update_example" {
  value = dnacenter_network_device_management_address_update.example
}
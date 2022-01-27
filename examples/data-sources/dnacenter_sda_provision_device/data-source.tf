
data "dnacenter_sda_provision_device" "example" {
  provider                     = dnacenter
  device_management_ip_address = "string"
}

output "dnacenter_sda_provision_device_example" {
  value = data.dnacenter_sda_provision_device.example.item
}

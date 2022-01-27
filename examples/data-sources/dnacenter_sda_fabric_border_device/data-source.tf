
data "dnacenter_sda_fabric_border_device" "example" {
  provider                     = dnacenter
  device_management_ip_address = "string"
}

output "dnacenter_sda_fabric_border_device_example" {
  value = data.dnacenter_sda_fabric_border_device.example.item
}

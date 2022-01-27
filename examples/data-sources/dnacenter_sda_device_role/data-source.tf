
data "dnacenter_sda_device_role" "example" {
  provider                     = dnacenter
  device_management_ip_address = "string"
}

output "dnacenter_sda_device_role_example" {
  value = data.dnacenter_sda_device_role.example.item
}

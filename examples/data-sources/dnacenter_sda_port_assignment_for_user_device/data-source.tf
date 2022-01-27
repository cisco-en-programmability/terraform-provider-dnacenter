
data "dnacenter_sda_port_assignment_for_user_device" "example" {
  provider                     = dnacenter
  device_management_ip_address = "string"
  interface_name               = "string"
}

output "dnacenter_sda_port_assignment_for_user_device_example" {
  value = data.dnacenter_sda_port_assignment_for_user_device.example.item
}

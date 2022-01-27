
data "dnacenter_sda_fabric_control_plane_device" "example" {
  provider                     = dnacenter
  device_management_ip_address = "string"
}

output "dnacenter_sda_fabric_control_plane_device_example" {
  value = data.dnacenter_sda_fabric_control_plane_device.example.item
}

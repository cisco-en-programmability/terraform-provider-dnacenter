
resource "dnacenter_sda_fabric_control_plane_device" "example" {
  provider = dnacenter
  parameters {

    device_management_ip_address = "string"
    site_name_hierarchy          = "string"
  }
}

output "dnacenter_sda_fabric_control_plane_device_example" {
  value = dnacenter_sda_fabric_control_plane_device.example
}
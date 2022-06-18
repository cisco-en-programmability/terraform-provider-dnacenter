
resource "dnacenter_sda_port_assignment_for_user_device" "example" {
  provider = dnacenter
  parameters {

    authenticate_template_name   = "string"
    data_ip_address_pool_name    = "string"
    device_management_ip_address = "string"
    interface_description        = "string"
    interface_name               = "string"
    scalable_group_name          = "string"
    site_name_hierarchy          = "string"
    voice_ip_address_pool_name   = "string"
  }
}

output "dnacenter_sda_port_assignment_for_user_device_example" {
  value = dnacenter_sda_port_assignment_for_user_device.example
}
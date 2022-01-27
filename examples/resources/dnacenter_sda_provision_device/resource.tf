
resource "dnacenter_sda_provision_device" "example" {
  provider = dnacenter
  parameters {

    device_management_ip_address = "string"
    site_name_hierarchy          = "string"
  }
}

output "dnacenter_sda_provision_device_example" {
  value = dnacenter_sda_provision_device.example
}
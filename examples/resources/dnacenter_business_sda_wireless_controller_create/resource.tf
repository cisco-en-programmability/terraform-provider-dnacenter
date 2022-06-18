
resource "dnacenter_business_sda_wireless_controller_create" "example" {
  provider = dnacenter
  parameters {

    device_name         = "string"
    site_name_hierarchy = "string"
  }
}

output "dnacenter_business_sda_wireless_controller_create_example" {
  value = dnacenter_business_sda_wireless_controller_create.example
}
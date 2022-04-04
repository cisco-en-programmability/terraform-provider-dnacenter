
resource "dnacenter_business_sda_wireless_controller_create" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    device_name         = "string"
    site_name_hierarchy = "string"
  }
}
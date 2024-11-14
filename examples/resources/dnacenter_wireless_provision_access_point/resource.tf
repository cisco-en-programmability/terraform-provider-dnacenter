
resource "dnacenter_wireless_provision_access_point" "example" {
  provider = dnacenter
  parameters {

    custom_ap_group_name   = "string"
    custom_flex_group_name = ["string"]
    device_name            = "string"
    rf_profile             = "string"
    site_name_hierarchy    = "string"
    type                   = "string"
  }
}

output "dnacenter_wireless_provision_access_point_example" {
  value = dnacenter_wireless_provision_access_point.example
}
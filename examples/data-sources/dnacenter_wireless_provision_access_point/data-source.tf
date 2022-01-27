
data "dnacenter_wireless_provision_access_point" "example" {
  provider               = dnacenter
  custom_ap_group_name   = "string"
  custom_flex_group_name = ["string"]
  device_name            = "string"
  rf_profile             = "string"
  site_id                = "string"
  site_name_hierarchy    = "string"
  type                   = "string"
}
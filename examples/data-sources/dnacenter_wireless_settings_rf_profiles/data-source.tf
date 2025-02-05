
data "dnacenter_wireless_settings_rf_profiles" "example" {
  provider                = dnacenter
  enable_radio_type6_g_hz = "false"
  enable_radio_type_a     = "false"
  enable_radio_type_b     = "false"
  limit                   = 1
  offset                  = 1
  rf_profile_name         = "string"
}

output "dnacenter_wireless_settings_rf_profiles_example" {
  value = data.dnacenter_wireless_settings_rf_profiles.example.items
}

data "dnacenter_wireless_settings_rf_profiles" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_wireless_settings_rf_profiles_example" {
  value = data.dnacenter_wireless_settings_rf_profiles.example.item
}

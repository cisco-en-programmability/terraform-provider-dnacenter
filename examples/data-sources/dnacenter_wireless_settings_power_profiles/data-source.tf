
data "dnacenter_wireless_settings_power_profiles" "example" {
  provider     = dnacenter
  limit        = 1
  offset       = 1
  profile_name = "string"
}

output "dnacenter_wireless_settings_power_profiles_example" {
  value = data.dnacenter_wireless_settings_power_profiles.example.items
}

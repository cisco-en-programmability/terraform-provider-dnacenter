
data "dnacenter_wireless_settings_ap_profiles" "example" {
  provider        = dnacenter
  ap_profile_name = "string"
  limit           = "string"
  offset          = "string"
}

output "dnacenter_wireless_settings_ap_profiles_example" {
  value = data.dnacenter_wireless_settings_ap_profiles.example.items
}


data "dnacenter_wireless_settings_power_profiles_id" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_wireless_settings_power_profiles_id_example" {
  value = data.dnacenter_wireless_settings_power_profiles_id.example.item
}

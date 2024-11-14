
data "dnacenter_wireless_settings_dot11be_profiles_count" "example" {
  provider = dnacenter
}

output "dnacenter_wireless_settings_dot11be_profiles_count_example" {
  value = data.dnacenter_wireless_settings_dot11be_profiles_count.example.item
}

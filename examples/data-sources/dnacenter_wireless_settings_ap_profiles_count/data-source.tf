
data "dnacenter_wireless_settings_ap_profiles_count" "example" {
  provider = dnacenter
}

output "dnacenter_wireless_settings_ap_profiles_count_example" {
  value = data.dnacenter_wireless_settings_ap_profiles_count.example.item
}

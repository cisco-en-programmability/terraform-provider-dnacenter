
data "dnacenter_wireless_settings_dot11be_profiles" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
}

output "dnacenter_wireless_settings_dot11be_profiles_example" {
  value = data.dnacenter_wireless_settings_dot11be_profiles.example.items
}

data "dnacenter_wireless_settings_dot11be_profiles" "example" {
  provider = dnacenter
  id       = "string"
}

output "dnacenter_wireless_settings_dot11be_profiles_example" {
  value = data.dnacenter_wireless_settings_dot11be_profiles.example.item
}

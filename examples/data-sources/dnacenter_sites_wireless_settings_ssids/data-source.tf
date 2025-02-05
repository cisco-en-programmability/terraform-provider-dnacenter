
data "dnacenter_sites_wireless_settings_ssids" "example" {
  provider    = dnacenter
  auth_type   = "string"
  l3auth_type = "string"
  limit       = 1
  offset      = 1
  site_id     = "string"
  ssid        = "string"
  wlan_type   = "string"
}

output "dnacenter_sites_wireless_settings_ssids_example" {
  value = data.dnacenter_sites_wireless_settings_ssids.example.items
}

data "dnacenter_sites_wireless_settings_ssids" "example" {
  provider    = dnacenter
  auth_type   = "string"
  l3auth_type = "string"
  limit       = 1
  offset      = 1
  site_id     = "string"
  ssid        = "string"
  wlan_type   = "string"
}

output "dnacenter_sites_wireless_settings_ssids_example" {
  value = data.dnacenter_sites_wireless_settings_ssids.example.item
}


data "dnacenter_sites_wireless_settings_ssids" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
  site_id  = "string"
}

output "dnacenter_sites_wireless_settings_ssids_example" {
  value = data.dnacenter_sites_wireless_settings_ssids.example.items
}

data "dnacenter_sites_wireless_settings_ssids" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
  site_id  = "string"
}

output "dnacenter_sites_wireless_settings_ssids_example" {
  value = data.dnacenter_sites_wireless_settings_ssids.example.item
}

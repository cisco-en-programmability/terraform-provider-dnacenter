
data "dnacenter_wireless_settings_ssids_override_at_sites" "example" {
  provider = dnacenter
  limit    = 1
  offset   = 1
  site_id  = "string"
}

output "dnacenter_wireless_settings_ssids_override_at_sites_example" {
  value = data.dnacenter_wireless_settings_ssids_override_at_sites.example.items
}

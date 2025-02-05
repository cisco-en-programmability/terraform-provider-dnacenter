
data "dnacenter_sites_wireless_settings_ssids_count" "example" {
  provider  = dnacenter
  inherited = "false"
  site_id   = "string"
}

output "dnacenter_sites_wireless_settings_ssids_count_example" {
  value = data.dnacenter_sites_wireless_settings_ssids_count.example.item
}

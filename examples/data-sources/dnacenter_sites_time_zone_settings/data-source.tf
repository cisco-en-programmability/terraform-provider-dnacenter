
data "dnacenter_sites_time_zone_settings" "example" {
  provider  = dnacenter
  id        = "string"
  inherited = "false"
}

output "dnacenter_sites_time_zone_settings_example" {
  value = data.dnacenter_sites_time_zone_settings.example.item
}

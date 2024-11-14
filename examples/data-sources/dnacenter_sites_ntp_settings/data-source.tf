
data "dnacenter_sites_ntp_settings" "example" {
  provider  = dnacenter
  id        = "string"
  inherited = "false"
}

output "dnacenter_sites_ntp_settings_example" {
  value = data.dnacenter_sites_ntp_settings.example.item
}

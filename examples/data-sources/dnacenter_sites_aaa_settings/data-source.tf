
data "dnacenter_sites_aaa_settings" "example" {
  provider  = dnacenter
  id        = "string"
  inherited = "false"
}

output "dnacenter_sites_aaa_settings_example" {
  value = data.dnacenter_sites_aaa_settings.example.item
}

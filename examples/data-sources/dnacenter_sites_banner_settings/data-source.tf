
data "dnacenter_sites_banner_settings" "example" {
  provider  = dnacenter
  id        = "string"
  inherited = "false"
}

output "dnacenter_sites_banner_settings_example" {
  value = data.dnacenter_sites_banner_settings.example.item
}

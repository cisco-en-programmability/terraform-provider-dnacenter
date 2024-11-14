
data "dnacenter_sites_dns_settings" "example" {
  provider  = dnacenter
  id        = "string"
  inherited = "false"
}

output "dnacenter_sites_dns_settings_example" {
  value = data.dnacenter_sites_dns_settings.example.item
}

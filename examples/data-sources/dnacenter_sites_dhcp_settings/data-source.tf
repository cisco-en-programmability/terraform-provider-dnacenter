
data "dnacenter_sites_dhcp_settings" "example" {
  provider  = dnacenter
  id        = "string"
  inherited = "false"
}

output "dnacenter_sites_dhcp_settings_example" {
  value = data.dnacenter_sites_dhcp_settings.example.item
}

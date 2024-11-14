
data "dnacenter_sites_telemetry_settings" "example" {
  provider  = dnacenter
  id        = "string"
  inherited = "false"
}

output "dnacenter_sites_telemetry_settings_example" {
  value = data.dnacenter_sites_telemetry_settings.example.item
}

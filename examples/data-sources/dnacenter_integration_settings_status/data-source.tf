
data "dnacenter_integration_settings_status" "example" {
  provider = dnacenter
}

output "dnacenter_integration_settings_status_example" {
  value = data.dnacenter_integration_settings_status.example.items
}

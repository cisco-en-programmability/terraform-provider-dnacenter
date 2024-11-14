
data "dnacenter_integration_settings_itsm_instances" "example" {
  provider = dnacenter
}

output "dnacenter_integration_settings_itsm_instances_example" {
  value = data.dnacenter_integration_settings_itsm_instances.example.items
}

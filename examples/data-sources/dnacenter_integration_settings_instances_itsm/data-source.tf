
data "dnacenter_integration_settings_instances_itsm" "example" {
  provider    = dnacenter
  instance_id = "string"
}

output "dnacenter_integration_settings_instances_itsm_example" {
  value = data.dnacenter_integration_settings_instances_itsm.example.item
}

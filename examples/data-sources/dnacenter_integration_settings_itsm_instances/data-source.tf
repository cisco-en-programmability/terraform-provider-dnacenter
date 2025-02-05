
data "dnacenter_integration_settings_itsm_instances" "example" {
  provider  = dnacenter
  order     = "string"
  page      = 1.0
  page_size = 1.0
  sort_by   = "string"
}

output "dnacenter_integration_settings_itsm_instances_example" {
  value = data.dnacenter_integration_settings_itsm_instances.example.item
}

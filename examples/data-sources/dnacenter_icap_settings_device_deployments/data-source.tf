
data "dnacenter_icap_settings_device_deployments" "example" {
  provider           = dnacenter
  deploy_activity_id = "string"
  limit              = 1
  network_device_ids = "string"
  offset             = 1
  order              = "string"
  sort_by            = "string"
}

output "dnacenter_icap_settings_device_deployments_example" {
  value = data.dnacenter_icap_settings_device_deployments.example.items
}

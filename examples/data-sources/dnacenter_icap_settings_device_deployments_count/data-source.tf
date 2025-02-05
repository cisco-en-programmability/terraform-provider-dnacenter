
data "dnacenter_icap_settings_device_deployments_count" "example" {
  provider           = dnacenter
  deploy_activity_id = "string"
  network_device_ids = "string"
}

output "dnacenter_icap_settings_device_deployments_count_example" {
  value = data.dnacenter_icap_settings_device_deployments_count.example.item
}


data "dnacenter_icap_settings_configuration_models_preview_activity_id_network_device_status_details" "example" {
  provider            = dnacenter
  preview_activity_id = "string"
}

output "dnacenter_icap_settings_configuration_models_preview_activity_id_network_device_status_details_example" {
  value = data.dnacenter_icap_settings_configuration_models_preview_activity_id_network_device_status_details.example.items
}

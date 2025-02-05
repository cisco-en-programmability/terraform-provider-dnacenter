
data "dnacenter_icap_settings_configuration_models_preview_activity_id_network_devices_network_device_id_config" "example" {
  provider            = dnacenter
  network_device_id   = "string"
  preview_activity_id = "string"
}

output "dnacenter_icap_settings_configuration_models_preview_activity_id_network_devices_network_device_id_config_example" {
  value = data.dnacenter_icap_settings_configuration_models_preview_activity_id_network_devices_network_device_id_config.example.item
}

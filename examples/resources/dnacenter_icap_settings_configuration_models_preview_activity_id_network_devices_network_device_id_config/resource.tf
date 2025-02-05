
resource "dnacenter_icap_settings_configuration_models_preview_activity_id_network_devices_network_device_id_config" "example" {
  provider = dnacenter
  parameters {

    network_device_id   = "string"
    object              = "string"
    preview_activity_id = "string"
  }
}

output "dnacenter_icap_settings_configuration_models_preview_activity_id_network_devices_network_device_id_config_example" {
  value = dnacenter_icap_settings_configuration_models_preview_activity_id_network_devices_network_device_id_config.example
}

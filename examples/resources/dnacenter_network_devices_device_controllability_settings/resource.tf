
resource "dnacenter_network_devices_device_controllability_settings" "example" {
  provider = dnacenter

  parameters {

    autocorrect_telemetry_config = "false"
    device_controllability       = "false"
  }
}

output "dnacenter_network_devices_device_controllability_settings_example" {
  value = dnacenter_network_devices_device_controllability_settings.example
}
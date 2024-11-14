
resource "dnacenter_network_devices_resync_interval_settings_override" "example" {
  provider = dnacenter
}

output "dnacenter_network_devices_resync_interval_settings_override_example" {
  value = dnacenter_network_devices_resync_interval_settings_override.example
}
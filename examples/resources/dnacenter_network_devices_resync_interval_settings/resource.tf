
resource "dnacenter_network_devices_resync_interval_settings" "example" {
  provider = dnacenter
  parameters {

    interval = 1
  }
}

output "dnacenter_network_devices_resync_interval_settings_example" {
  value = dnacenter_network_devices_resync_interval_settings.example
}